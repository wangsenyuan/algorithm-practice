package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) []int {
	var x []int
	var y []int
	for i, v := range a {
		if i&1 == 0 {
			x = append(x, v)
		} else {
			y = append(y, v)
		}
	}
	flag := (countInversions(x) & 1) != (countInversions(y) & 1)
	sort.Ints(x)
	sort.Ints(y)
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i&1 == 0 {
			res[i] = x[0]
			x = x[1:]
		} else {
			res[i] = y[0]
			y = y[1:]
		}
	}
	n := len(a)
	if flag {
		res[n-3], res[n-1] = res[n-1], res[n-3]
	}
	return res
}

func countInversions(arr []int) int {
	n := len(arr)*2 + 1
	bit := make(BIT, n+3)
	var res int
	for _, x := range arr {
		res += bit.Get(n) - bit.Get(x)
		bit.Set(x, 1)
	}
	return res
}

type BIT []int

func (bit BIT) Set(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}
func (bit BIT) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func solve1(a []int) []int {
	n := len(a)
	// 3, 4, 1, 2, 5 => 1, 2, 3, 4, 5
	// 那么把1排在第一位后，2也被带到了第二位，它的后面，就要变成3
	// 还有一个就是原来4的后面是1，现在要变成5
	// 这个过程相当于，把1，2同时移动到了前面
	// 一开始，最后一个是作不能作为为首位去移动的
	// 但是它是作为第二位，有可能被移动到前面去。然后就可以移动了
	// 少于4个时候，是没法移动的

	pqs := make([]IntHeap, 2)
	prev := make([]int, n+1)
	next := make([]int, n+2)
	for i := range a {
		heap.Push(&pqs[i&1], a[i])
		if i > 0 {
			prev[a[i]] = a[i-1]
		}
		if i+1 < n {
			next[a[i]] = a[i+1]
		} else {
			next[a[i]] = n + 1
		}
	}
	next[n+1] = n + 1

	res := make([]int, 0, n)

	head := a[0]

	for i := range n {
		if len(pqs[0])+len(pqs[1]) == 3 {
			break
		}
		res = append(res, heap.Pop(&pqs[i&1]).(int))
		u := res[i]

		if u == head {
			head = next[head]
			prev[head] = 0
			continue
		}

		if next[u] > n {
			// head, x...y, p, u => head, p, u, x.... y
			x := next[head]
			p := prev[u]
			y := prev[p]
			next[head] = p
			prev[p] = head
			next[u] = x
			prev[x] = u
			next[y] = n + 1
		}
		// u != head
		p := prev[u]
		v := next[u]
		w := next[v]
		// head, ....p, u, v, w
		next[p] = w
		if w <= n {
			prev[w] = p
		}
		next[v] = head
		prev[head] = v
		prev[v] = 0
		head = v
	}

	for head <= n {
		res = append(res, head)
		head = next[head]
	}

	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
