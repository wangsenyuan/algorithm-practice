package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscan(reader, &m)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, a, queries)
}

func solve(n int, a []int, queries [][]int) []int {
	pos := make([]int, n)
	for i := range a {
		a[i]--
		pos[a[i]] = i
	}

	swap := func(x int, y int) {
		a[x], a[y] = a[y], a[x]
		pos[a[x]] = x
		pos[a[y]] = y
	}

	tr := Build(pos)

	var ans []int

	for _, cur := range queries {
		x, y := cur[1], cur[2]
		x--
		y--
		if cur[0] == 1 {
			ans = append(ans, tr.Get(x, y)+1)
		} else {
			swap(x, y)
			tr.Update(a[x], x)
			tr.Update(a[y], y)
		}
	}

	return ans
}

type Tree struct {
	val [][2]int
	cnt []int
	n   int
}

func (tr *Tree) merge(i int) {
	tr.cnt[i] = tr.cnt[i*2+1] + tr.cnt[i*2+2]
	if tr.val[i*2+1][1] > tr.val[i*2+2][0] {
		tr.cnt[i]++
	}
	tr.val[i][0] = tr.val[i*2+1][0]
	tr.val[i][1] = tr.val[i*2+2][1]
}

func Build(arr []int) *Tree {
	n := len(arr)
	tr := new(Tree)
	tr.n = n
	tr.val = make([][2]int, 4*n)
	tr.cnt = make([]int, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = [...]int{arr[l], arr[l]}
			tr.cnt[i] = 0
			return
		}
		mid := (l + r) / 2
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		tr.merge(i)
	}
	f(0, 0, n-1)
	return tr
}

func (tr *Tree) Update(p int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.val[i] = [...]int{v, v}
			tr.cnt[i] = 0
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr.merge(i)
	}
	f(0, 0, tr.n-1)
}

func (tr *Tree) Get(L int, R int) int {
	var f func(i int, l int, r int, L int, R int) (int, [2]int)
	f = func(i int, l int, r int, L int, R int) (cnt int, val [2]int) {
		if L == l && R == r {
			return tr.cnt[i], tr.val[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return f(i*2+1, l, mid, L, R)
		}
		if mid < L {
			return f(i*2+2, mid+1, r, L, R)
		}
		c1, v1 := f(i*2+1, l, mid, L, mid)
		c2, v2 := f(i*2+2, mid+1, r, mid+1, R)
		cnt = c1 + c2
		if v1[1] > v2[0] {
			cnt++
		}
		val[0] = v1[0]
		val[1] = v2[1]
		return
	}
	cnt, _ := f(0, 0, tr.n-1, L, R)
	return cnt
}
