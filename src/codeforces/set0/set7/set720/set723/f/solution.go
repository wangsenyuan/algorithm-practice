package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, _, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Fprintln(writer, "No")
		return
	}
	fmt.Fprintln(writer, "Yes")
	for _, e := range res {
		fmt.Fprintln(writer, e[0], e[1])
	}
}

func drive(reader *bufio.Reader) (n int, s int, t int, ds int, dt int, ok bool, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	fmt.Fscan(reader, &s, &t, &ds, &dt)
	ok, res = solve(n, edges, s, t, ds, dt)
	return

}

func solve(n int, edges [][]int, s int, t int, ds int, dt int) (ok bool, res [][]int) {

	s--
	t--

	if s > t {
		s, t = t, s
		ds, dt = dt, ds
	}

	var g1 [][]int
	var g2 [][]int
	var g3 [][]int

	set := NewDSU(n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		u, v = min(u, v), max(u, v)
		if u == s && v == t {
			// 这个只会有一条
			g3 = append(g3, e)
		} else if u == s || v == s {
			g1 = append(g1, e)
		} else if v == t || u == t {
			g2 = append(g2, e)
		} else if set.Union(u, v) {
			res = append(res, e)
		}
	}

	// 剩下的都可以当作点来处理
	flag := make([]int, n)

	for _, e := range g1 {
		u, v := e[0]-1, e[1]-1
		w := u ^ v ^ s
		w = set.Find(w)
		flag[w] |= 1
	}

	for _, e := range g2 {
		u, v := e[0]-1, e[1]-1
		w := u ^ v ^ t
		w = set.Find(w)
		flag[w] |= 2
	}
	// 然后找出那些只能和s相连的部分

	add := func(e []int, d *int) {
		u, v := e[0]-1, e[1]-1
		if set.Union(u, v) {
			*d--
			res = append(res, e)
		}
	}

	g4 := make([][][]int, n)
	var arr []int

	for _, e := range g1 {
		u, v := e[0]-1, e[1]-1
		w := u ^ v ^ s
		w = set.Find(w)

		if flag[w] == 3 {
			// 这个既可以连接到s, 也可以连接到t
			g4[w] = append(g4[w], e)
			arr = append(arr, w)
			continue
		}
		add(e, &ds)
	}

	g5 := make([][][]int, n)

	for _, e := range g2 {
		u, v := e[0]-1, e[1]-1
		w := u ^ v ^ t
		w = set.Find(w)

		if flag[w] == 3 {
			g5[w] = append(g5[w], e)
			arr = append(arr, w)
			continue
		}

		add(e, &dt)
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)

	if ds <= 0 || dt <= 0 {
		return false, nil
	}

	for i := 0; i < len(arr); i++ {
		w := arr[i]
		if i == 0 {
			add(g4[w][0], &ds)
			add(g5[w][0], &dt)
		} else {
			if ds >= dt {
				add(g4[w][0], &ds)
			} else {
				add(g5[w][0], &dt)
			}
		}
	}

	if set.Find(s) != set.Find(t) {
		if len(g3) == 0 {
			return false, nil
		}
		add(g3[0], &ds)
		dt--
	}

	if ds < 0 || dt < 0 {
		return false, nil
	}

	return true, res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}
