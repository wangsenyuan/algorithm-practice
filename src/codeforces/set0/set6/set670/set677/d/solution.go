package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, p int
	fmt.Fscan(reader, &n, &m, &p)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(p, a)
}

func solve(p int, a [][]int) int {
	if p == 1 {
		// 只有一个格子
		return 0
	}
	n := len(a)
	m := len(a[0])

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
		for j := range m {
			dist[i][j] = inf
		}
	}

	// dist[0][0] = 0
	// p <= n * m, 没法进行 n * m * p
	layers := make([][]int, p)
	for i := range n {
		for j := range m {
			x := a[i][j] - 1
			layers[x] = append(layers[x], i*m+j)
		}
	}

	for _, cur := range layers[0] {
		i, j := cur/m, cur%m
		dist[i][j] = i + j
	}

	w := max(n, m)
	t1 := NewSegTree(w)
	t2 := NewSegTree(w)

	reset := func(arr []int) {
		for j := range len(arr) {
			c := arr[j] % m

			t1.Update(c, inf)
			t2.Update(c, inf)
		}
	}

	for k := 1; k < p; k++ {
		prev := layers[k-1]
		cur := layers[k]
		// 要在当前(x, y)距离内找到，最近的prev的的点
		// 如果是它左边的点, (i, j), (i1, j1)
		// 一开始按照行来处理
		for i, j := 0, 0; i < len(cur); i++ {
			for j < len(prev) && prev[j]/m <= cur[i]/m {
				r, c := prev[j]/m, prev[j]%m
				// 用于计算前半段
				t1.Update(c, dist[r][c]-r-c)
				// 用于计算后半段
				t2.Update(c, dist[r][c]-r+c)
				j++
			}
			r, c := cur[i]/m, cur[i]%m
			w1 := t1.Get(0, c) + r + c
			w2 := t2.Get(c, m) + r - c
			dist[r][c] = min(w1, w2)
		}
		reset(prev)
		for i, j := len(cur)-1, len(prev)-1; i >= 0; i-- {
			for j >= 0 && prev[j]/m >= cur[i]/m {
				r, c := prev[j]/m, prev[j]%m
				t1.Update(c, dist[r][c]+r-c)
				t2.Update(c, dist[r][c]+r+c)
				j--
			}
			r, c := cur[i]/m, cur[i]%m
			w1 := t1.Get(0, c) - r + c
			w2 := t2.Get(c, m) - r - c
			dist[r][c] = min(dist[r][c], w1, w2)
		}

		reset(prev)
	}

	pr, pc := layers[p-1][0]/m, layers[p-1][0]%m
	return dist[pr][pc]
}

type SegTree []int

const inf = 1 << 30

func NewSegTree(n int) SegTree {
	arr := make(SegTree, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v
	for p > 1 {
		tr[p>>1] = min(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
