package p3873

import (
	"cmp"
	"slices"
)

func maxActivated(points [][]int) int {
	n := len(points)
	set := NewDSU(n)
	slices.SortFunc(points, func(a []int, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	last := make(map[int]int)

	for i := 0; i < n; {
		j := i
		for i < n && points[i][0] == points[j][0] {
			if i > j {
				set.Union(i, j)
			}
			if k, ok := last[points[i][1]]; ok {
				set.Union(i, k)
			}
			last[points[i][1]] = i
			i++
		}
	}
	var sz []int
	for i := range n {
		if i == set.Find(i) {
			sz = append(sz, set.cnt[i])
		}
	}
	slices.Sort(sz)
	slices.Reverse(sz)
	if len(sz) == 1 {
		return n + 1
	}
	return sz[0] + sz[1] + 1
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
