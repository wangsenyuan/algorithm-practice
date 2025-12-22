package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const H = 32

type Node struct {
	cnt      int
	children [2]*Node
}

func solve(A []int, k int) int {
	n := len(A)
	a := make([]int, n+1)
	copy(a[1:], A)

	root := new(Node)

	add := func(num int) {
		node := root
		for i := H - 1; i >= 0; i-- {
			x := (num >> i) & 1
			if node.children[x] == nil {
				node.children[x] = new(Node)
			}
			node = node.children[x]
			node.cnt++
		}
	}

	get := func(num int, k int) int {
		var res int
		node := root
		for i := H - 1; i >= 0; i-- {
			x := (num >> i) & 1
			if node.children[x^1] != nil {
				if k <= node.children[1^x].cnt {
					res |= 1 << i
					x ^= 1
				} else {
					k -= node.children[x^1].cnt
				}
			}

			node = node.children[x]
		}

		return res
	}

	add(0)
	for i := 1; i <= n; i++ {
		a[i] ^= a[i-1]
		add(a[i])
	}
	pq := make(PQ, 0, n+1)
	heap.Push(&pq, &Item{})
	for i := 0; i <= n; i++ {
		heap.Push(&pq, &Item{get(a[i], 1), i, 1})
	}

	var res int

	for k *= 2; k > 0; k-- {
		it := pq[0]
		res += it.xor
		i := it.id
		it.k++
		it.xor = get(a[i], it.k)
		heap.Fix(&pq, 0)
	}

	return res / 2
}

type Item struct {
	xor int
	id  int
	k   int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].xor > pq[j].xor
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	it := x.(*Item)
	*pq = append(*pq, it)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

func solve1(a []int, k int) int {
	// a[l] ^ ... ^ a[r] =
	// sum of [1...k]
	// 感觉还是要贡献法
	// 从最高位开始,d, 如果它能贡献的数量 >= k, 那么最后的结果中 (1 << d) * k
	// 在低位处理的时候，需要保证高位保持
	// 如果数量 < k, 那么应该分成两类，一类是是保证d设置了值，假设有x个（可以把x当成新的k）,另外一类是 k - x

	n := len(a)

	var dfs func(d int, mask int, k int) int

	dfs = func(d int, mask int, k int) int {
		// 现在需要知道在保证mask的前提下，有多少个异或区间 = 1
		if d < 0 || k == 0 {
			return 0
		}
		ma := mask << 1
		mb := ma | 1
		freq := make(map[int]int)
		cnt := make([]int, 2)
		var sum int
		freq[0] = 1
		for i := 0; i < n; i++ {
			num := a[i] >> d
			sum ^= num
			cnt[0] += freq[sum^ma]
			cnt[1] += freq[sum^mb]
			freq[sum]++
		}
		if cnt[1] >= k {
			return (1<<d)*k + dfs(d-1, mb, k)
		}
		res := (1<<d)*cnt[1] + dfs(d-1, mb, cnt[1])
		// 这里有问题，在保证ma的情况下，会重复计算，而且它们
		res += dfs(d-1, ma, k-cnt[1])
		return res
	}

	x := a[0]
	for i := 0; i < n; i++ {
		if x < a[i] {
			x = a[i]
		}
	}

	cnt := bits.LeadingZeros64(uint64(x))

	return dfs(64-cnt, 0, k)
}

func bruteForce(arr []int, k int) int {
	var sum []int
	for l := 0; l < len(arr); l++ {
		var s int
		for r := l; r < len(arr); r++ {
			s ^= arr[r]
			sum = append(sum, s)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sum)))

	var res int
	for i := 0; i < k; i++ {
		res += sum[i]
	}
	return res
}
