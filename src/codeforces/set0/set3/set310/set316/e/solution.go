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
	fmt.Print(buf.String())
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
	queries := make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '3' {
			queries[i] = make([]int, 4)
		} else {
			queries[i] = make([]int, 3)
		}
		var pos int
		for j := range queries[i] {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(a, queries)
}

const mod = 1000000000

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a, b int) int {
	return a * b % mod
}

const N = 2000010

var F [N]int
var SF [N]int

func init() {
	F[0] = 0
	F[1] = 1
	for i := 2; i < N; i++ {
		F[i] = add(F[i-1], F[i-2])
	}
	SF[0] = 0
	for i := 1; i < N; i++ {
		SF[i] = add(F[i], SF[i-1])
	}
}

func solve(a []int, queries [][]int) []int {
	tr := NewTree(a)

	var ans []int
	for _, cur := range queries {
		if cur[0] == 1 {
			x, v := cur[1], cur[2]
			x--
			tr.Set(x, v)
		} else if cur[0] == 3 {
			l, r, d := cur[1]-1, cur[2]-1, cur[3]
			tr.Update(l, r, d)
		} else {
			l, r := cur[1]-1, cur[2]-1
			ans = append(ans, tr.Query(l, r))
		}
	}
	return ans
}

type node struct {
	s1 int
	s2 int
	sz int
}

func merge(x node, y node) node {
	if x.sz == 0 {
		return y
	}
	var c node
	c.sz = x.sz + y.sz
	c.s1 = add(x.s1, mul(y.s1, F[x.sz+1]), mul(y.s2, F[x.sz]))
	c.s2 = add(x.s2, mul(y.s1, F[x.sz]), mul(y.s2, F[x.sz-1]))
	return c
}

type Tree struct {
	val  []node
	lazy []int
	sz   int
}

func NewTree(a []int) *Tree {
	n := len(a)
	val := make([]node, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val[i] = node{a[l], 0, 1}
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		val[i] = merge(val[i*2+1], val[i*2+2])
	}
	build(0, 0, n-1)
	return &Tree{val, make([]int, 4*n), n}
}

func (tr *Tree) update(i int, w int) {
	tr.lazy[i] = add(tr.lazy[i], w)
	tr.val[i].s1 = add(tr.val[i].s1, mul(SF[tr.val[i].sz], w))
	tr.val[i].s2 = add(tr.val[i].s2, mul(SF[tr.val[i].sz-1], w))
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] != 0 {
		tr.update(2*i+1, tr.lazy[i])
		tr.update(2*i+2, tr.lazy[i])
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Set(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = node{v, 0, 1}
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.val[i] = merge(tr.val[i*2+1], tr.val[i*2+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Update(L int, R int, d int) {
	var loop func(i int, l int, r int, L int, R int)
	loop = func(i int, l int, r int, L int, R int) {
		if L == l && r == R {
			tr.update(i, d)
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if L <= mid {
			loop(2*i+1, l, mid, L, min(R, mid))
		}
		if mid < R {
			loop(2*i+2, mid+1, r, max(L, mid+1), R)
		}
		tr.val[i] = merge(tr.val[2*i+1], tr.val[2*i+2])
	}
	loop(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Query(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) node
	loop = func(i int, l int, r int, L int, R int) node {
		if L == l && R == r {
			return tr.val[i]
		}
		tr.push(i)
		var res node
		mid := (l + r) / 2
		if L <= mid {
			res = loop(2*i+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res = merge(res, loop(2*i+2, mid+1, r, max(L, mid+1), R))
		}
		return res
	}
	res := loop(0, 0, tr.sz-1, L, R)
	return res.s1
}
