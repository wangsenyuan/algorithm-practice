package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, ln := readTwoNums(reader)
	a := readNNums(reader, n)
	k := readNum(reader)
	return solve(k, ln, a)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(k int, ln int, a []int) int {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		if x.first == y.first {
			return x.second - y.second
		}
		return x.first - y.first
	})

	findPos := func(x pair) int {
		return sort.Search(n, func(i int) bool {
			return arr[i].first > x.first || arr[i].first == x.first && arr[i].second >= x.second
		})
	}

	res := -inf

	tree := NewTree(n)
	cnt := make([]int, 2)
	var sum int
	for i := range n {
		sum += a[i]
		if a[i] >= 0 {
			cnt[1]++
		} else {
			cnt[0]++
		}
		j := findPos(pair{a[i], i})
		tree.Update(j, a[i], true)
		if i >= ln {
			k := findPos(pair{a[i-ln], i - ln})
			tree.Update(k, a[i-ln], false)
			sum -= a[i-ln]
			if a[i-ln] >= 0 {
				cnt[1]--
			} else {
				cnt[0]--
			}
		}
		if i >= ln-1 {
			// 不改变
			res = max(res, abs(sum))
			// 先改变负数
			x := min(cnt[0], k)
			if x > 0 {
				tmp := tree.GetMin(x)
				res = max(res, abs(sum-2*tmp))
			}
			x = min(cnt[1], k)
			if x > 0 {
				tmp := tree.GetMax(x)
				res = max(res, abs(sum-2*tmp))
			}
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}

type Tree struct {
	sum []int
	cnt []int
	sz  int
}

func NewTree(n int) *Tree {
	sum := make([]int, n*4)
	cnt := make([]int, n*4)
	return &Tree{sum, cnt, n}
}

func (t *Tree) Update(p int, v int, add bool) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			if add {
				t.sum[i] += v
				t.cnt[i]++
			} else {
				t.sum[i] -= v
				t.cnt[i]--
			}
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		t.sum[i] = t.sum[i*2+1] + t.sum[i*2+2]
		t.cnt[i] = t.cnt[i*2+1] + t.cnt[i*2+2]
	}
	loop(0, 0, t.sz-1)
}

func (t *Tree) GetMin(x int) int {
	var loop func(i int, l int, r int, x int) int
	loop = func(i int, l int, r int, x int) int {
		if t.cnt[i] == x {
			return t.sum[i]
		}
		mid := (l + r) / 2
		res := loop(2*i+1, l, mid, min(x, t.cnt[2*i+1]))
		if x > t.cnt[2*i+1] {
			res += loop(2*i+2, mid+1, r, x-t.cnt[2*i+1])
		}
		return res
	}
	return loop(0, 0, t.sz-1, x)
}

func (t *Tree) GetMax(x int) int {
	var loop func(i int, l int, r int, x int) int
	loop = func(i int, l int, r int, x int) int {
		if t.cnt[i] == x {
			return t.sum[i]
		}
		mid := (l + r) / 2
		res := loop(2*i+2, mid+1, r, min(x, t.cnt[2*i+2]))
		if x > t.cnt[2*i+2] {
			res += loop(2*i+1, l, mid, x-t.cnt[2*i+2])
		}
		return res
	}
	return loop(0, 0, t.sz-1, x)
}
