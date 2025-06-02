package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	res := process(reader)
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
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	k := readNum(reader)
	queries := make([][]int, k)
	for i := range k {
		queries[i] = readNNums(reader, 3)
	}
	return solve(a, b, queries)
}

const X = 1000000

func solve(a []int, b []int, queries [][]int) []int {
	tr := NewTree(X + 1)
	for _, x := range a {
		tr.Update(0, x, 1)
	}
	for _, y := range b {
		tr.Update(0, y, -1)
	}
	ans := make([]int, len(queries))
	for j, cur := range queries {
		if cur[0] == 1 {
			i, x := cur[1]-1, cur[2]
			tr.Update(0, a[i], -1)
			a[i] = x
			tr.Update(0, x, 1)
		} else {
			i, x := cur[1]-1, cur[2]
			tr.Update(0, b[i], 1)
			b[i] = x
			tr.Update(0, b[i], -1)
		}
		ans[j] = tr.Find()
		if ans[j] == 0 {
			ans[j]--
		}
	}
	return ans
}

type Tree struct {
	val  []int
	lazy []int
	n    int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	return &Tree{val, lazy, n}
}

func (t *Tree) update(i int, l int, r int, v int) {
	t.val[i] += v
	t.lazy[i] += v
}

func (t *Tree) push(i int, l int, r int) {
	if t.lazy[i] != 0 {
		mid := (l + r) / 2
		t.update(2*i+1, l, mid, t.lazy[i])
		t.update(2*i+2, mid, r, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) pull(i int) {
	t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int, L int, R int)
	loop = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.update(i, l, r, v)
			return
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		if L <= mid {
			loop(2*i+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			loop(2*i+2, mid+1, r, max(L, mid+1), R)
		}
		t.pull(i)
	}

	loop(0, 0, t.n-1, L, R)
}

func (t *Tree) Find() int {
	var loop func(i int, l int, r int) int
	loop = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		t.push(i, l, r)
		mid := (l + r) / 2
		if t.val[2*i+2] > 0 {
			return loop(2*i+2, mid+1, r)
		}
		return loop(2*i+1, l, mid)
	}
	if t.val[0] == 0 {
		return -1
	}
	return loop(0, 0, t.n-1)
}
