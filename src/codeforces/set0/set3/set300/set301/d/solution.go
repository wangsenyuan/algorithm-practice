package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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
	n, m := readTwoNums(reader)
	p := readNNums(reader, n)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(p, qs)
}

type pair struct {
	first  int
	second int
}

func solve(p []int, queries [][]int) []int {

	n := len(p)
	pos := make([]int, n+1)

	for i := range n {
		pos[p[i]] = i
	}

	arr := make([][]int, n)

	for x := 1; x <= n; x++ {
		for y := x; y <= n; y += x {
			l := min(pos[x], pos[y])
			r := max(pos[x], pos[y])
			arr[r] = append(arr[r], l)
		}
	}

	at := make([][]pair, n)
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		l--
		r--
		at[r] = append(at[r], pair{l, i})
	}

	m := len(queries)

	res := make([]int, m)

	sum := NewSegTree(n)

	for r := 0; r < n; r++ {
		for _, l := range arr[r] {
			sum.Add(l, 1)
		}
		for _, cur := range at[r] {
			l := cur.first
			i := cur.second
			res[i] = sum.Query(l, r+1)
		}
	}

	return res
}

type SegTree struct {
	arr []int
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	return &SegTree{arr, n}
}

func (seg *SegTree) Add(p int, v int) {
	p += seg.sz
	seg.arr[p] += v
	for p > 1 {
		seg.arr[p>>1] = seg.arr[p] + seg.arr[p^1]
		p >>= 1
	}
}

func (seg *SegTree) Query(l int, r int) int {
	l += seg.sz
	r += seg.sz
	var res int
	for l < r {
		if l&1 == 1 {
			res += seg.arr[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += seg.arr[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}
