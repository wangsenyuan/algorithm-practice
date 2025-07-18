package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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
		if x == len(bs) {
			return res[:i]
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n, q, s := readThreeNums(reader)

	readPlan := func() []int {
		return readNNums(reader, 5)
	}

	return solve(n, s, q, readPlan)
}

type data struct {
	l int
	r int
	w int
}

func solve(n int, s int, q int, readPlan func() []int) []int {
	// for 1 and 2
	g := make([][]data, n)
	var q3 [][]int

	for range q {
		plan := readPlan()
		switch plan[0] {
		case 1:
			u, v, w := plan[1]-1, plan[2]-1, plan[3]
			g[u] = append(g[u], data{v, v, w})
		case 2:
			v, l, r, w := plan[1]-1, plan[2]-1, plan[3]-1, plan[4]
			g[v] = append(g[v], data{l, r, w})
		default:
			q3 = append(q3, plan)
		}
	}
	slices.SortFunc(q3, func(a, b []int) int {
		return b[3] - a[3]
	})

	g3 := make([][]data, n)
	for _, plan := range q3 {
		v, l, r, w := plan[1]-1, plan[2]-1, plan[3]-1, plan[4]
		g3[l] = append(g3[l], data{v, r, w})
	}

	val3 := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			val3[i] = -1
			if len(g3[l]) > 0 {
				val3[i] = g3[l][0].r
			}
			return
		}
		mid := (l + r) >> 1
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		val3[i] = max(val3[2*i+1], val3[2*i+2])
	}

	build(0, 0, n-1)

	var todo []data

	var findAndPop func(i int, l int, r int, v int)
	findAndPop = func(i int, l int, r int, v int) {
		if v < l || val3[i] < v {
			return
		}
		// l <= v and v <= val3[i]
		if l == r {
			for len(g3[l]) > 0 && g3[l][0].r >= v {
				todo = append(todo, g3[l][0])
				g3[l] = g3[l][1:]
			}

			val3[i] = -1
			if len(g3[l]) > 0 {
				val3[i] = g3[l][0].r
			}
			return
		}
		mid := (l + r) >> 1

		findAndPop(2*i+1, l, mid, v)
		findAndPop(2*i+2, mid+1, r, v)

		val3[i] = max(val3[2*i+1], val3[2*i+2])
	}

	tr := NewTree(n)
	tr.Update(s-1, s-1, 0)

	ans := make([]int, n)

	for i := range n {
		ans[i] = inf
	}

	for {
		tmp := tr.QueryMinPosition()
		u, du := tmp[0], tmp[1]
		if du == inf {
			break
		}
		ans[u] = du
		tr.Invalidate(u)

		for _, cur := range g[u] {
			l, r, w := cur.l, cur.r, cur.w
			tr.Update(l, r, du+w)
		}

		todo = todo[:0]
		findAndPop(0, 0, n-1, u)

		for _, cur := range todo {
			v, w := cur.l, cur.w
			tr.Update(v, v, du+w)
		}
	}

	for i := range n {
		if ans[i] == inf {
			ans[i] = -1
		}
	}

	return ans
}

const inf = 1 << 60

type Tree struct {
	val    []int
	lazy   []int
	marked []bool
	sz     int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	marked := make([]bool, 4*n)
	for i := range len(val) {
		val[i] = inf
		lazy[i] = inf
	}

	return &Tree{val, lazy, marked, n}
}

func (t *Tree) update(i int, v int) {
	if t.marked[i] {
		return
	}
	t.val[i] = min(t.val[i], v)
	t.lazy[i] = min(t.lazy[i], v)
}

func (t *Tree) push(i int) {
	t.update(2*i+1, t.lazy[i])
	t.update(2*i+2, t.lazy[i])
	t.lazy[i] = inf
}

func (t *Tree) Update(L int, R int, v int) {
	var loop func(i int, l int, r int, L int, R int)
	loop = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if R <= mid {
			loop(2*i+1, l, mid, L, R)
		} else if mid < L {
			loop(2*i+2, mid+1, r, L, R)
		} else {
			loop(2*i+1, l, mid, L, mid)
			loop(2*i+2, mid+1, r, mid+1, R)
		}
		t.val[i] = min(t.val[i*2+1], t.val[2*i+2])
	}
	loop(0, 0, t.sz-1, L, R)
}

func (t *Tree) Invalidate(p int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.marked[i] = true
			t.val[i] = inf
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		t.marked[i] = t.marked[2*i+1] && t.marked[2*i+2]
		if t.marked[i] {
			t.val[i] = inf
		} else {
			t.val[i] = min(t.val[i*2+1], t.val[2*i+2])
		}
	}
	loop(0, 0, t.sz-1)
}

func (t *Tree) QueryMinPosition() []int {
	if t.val[0] == inf {
		return []int{-1, inf}
	}

	var loop func(i int, l int, r int) []int
	loop = func(i int, l int, r int) []int {
		if l == r {
			return []int{l, t.val[i]}
		}
		t.push(i)
		mid := (l + r) >> 1
		if t.val[2*i+1] == t.val[i] {
			return loop(2*i+1, l, mid)
		}
		return loop(2*i+2, mid+1, r)
	}

	return loop(0, 0, t.sz-1)
}
