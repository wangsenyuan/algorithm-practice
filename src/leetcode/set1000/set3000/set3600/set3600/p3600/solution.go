package p3600

import "slices"

const inf = 1 << 30

func maxStability(n int, edges [][]int, k int) int {

	var mr int = inf
	var must [][]int
	var other [][]int
	set := NewDSU(n)

	for _, e := range edges {
		if e[3] == 1 {
			mr = min(mr, e[2])
			must = append(must, e)
			u, v := e[0], e[1]
			if set.Find(u) == set.Find(v) {
				return -1
			}
			set.Union(u, v)
		} else {
			other = append(other, e)
		}
	}

	slices.SortFunc(other, func(a, b []int) int {
		return b[2] - a[2]
	})

	check := func(exp int) bool {
		if exp > mr {
			return false
		}
		set.Reset()

		for _, e := range must {
			set.Union(e[0], e[1])
		}
		var i int
		for i < len(other) && other[i][2] >= exp {
			set.Union(other[i][0], other[i][1])
			i++
		}
		cnt := k
		for i < len(other) && cnt > 0 {
			u, v, w := other[i][0], other[i][1], other[i][2]
			if w*2 < exp {
				break
			}
			if set.Union(u, v) {
				cnt--
			}
			i++
		}

		return set.sz == 1
	}

	l, r := 0, mr+2
	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	r--
	if r > mr {
		return -1
	}
	return r
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

func (this *DSU) Reset() {
	for i := range this.arr {
		this.arr[i] = i
		this.cnt[i] = 1
	}
	this.sz = len(this.arr)
}
