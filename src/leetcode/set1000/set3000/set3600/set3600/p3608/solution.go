package p3608

import "slices"

func minTime(n int, edges [][]int, k int) int {
	slices.SortFunc(edges, func(a, b []int) int {
		return b[2] - a[2]
	})

	set := NewDSU(n)

	var ans int
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		set.Union(u, v)
		if set.sz < k {
			ans = w
			break
		}
	}

	return ans
}

type DSU struct {
	arr []int
	cnt []int
	sz  int
}

func NewDSU(n int) *DSU {
	d := new(DSU)
	d.arr = make([]int, n)
	d.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		d.arr[i] = i
		d.cnt[i]++
	}
	d.sz = n
	return d
}

func (d *DSU) Find(u int) int {
	if d.arr[u] != u {
		d.arr[u] = d.Find(d.arr[u])
	}
	return d.arr[u]
}

func (d *DSU) Union(a, b int) bool {
	a = d.Find(a)
	b = d.Find(b)
	if a == b {
		return false
	}
	if d.cnt[a] < d.cnt[b] {
		a, b = b, a
	}
	d.cnt[a] += d.cnt[b]
	d.arr[b] = a
	d.sz--
	return true
}
