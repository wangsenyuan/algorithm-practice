package p3613

import "slices"

func minCost(n int, edges [][]int, k int) int {
	if k == n {
		return 0
	}
	slices.SortFunc(edges, func(a, b []int) int {
		return a[2] - b[2]
	})


	set := NewDSU(n)

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		set.Union(u, v)
		if set.sz == k {
			return w
		}
	}
	return -1
}

type DSU struct {
	arr []int
	cnt []int
	sz  int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt, n}
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
	this.sz--
	return true
}
