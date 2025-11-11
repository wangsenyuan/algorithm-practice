package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect {
		return
	}

	type edge struct {
		u  int
		v  int
		w  int
		id int
	}

	arr := make([]edge, len(res))
	for i, cur := range res {
		arr[i] = edge{cur[0], cur[1], edges[i][0], i}
	}

	slices.SortFunc(arr, func(a, b edge) int {
		return a.w - b.w
	})

	set := NewDSU(n)

	for _, cur := range arr {
		u, v, id := cur.u, cur.v, cur.id
		u--
		v--
		if set.Find(u) != set.Find(v) {
			if edges[id][1] == 0 {
				t.Fatalf("Sample result %v not correct, it should not use edge %v, but it used", res, cur)
			}
			set.Union(u, v)
		} else {
			if edges[id][1] == 1 {
				t.Fatalf("Sample result %v not correct, it should use edge %v, but it did not use", res, cur)
			}
		}
	}

	root := set.Find(0)
	if set.cnt[root] != n {
		t.Fatalf("Sample result %v not correct, it should have one component with size %d, but it has %d", res, n, set.cnt[root])
	}
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

func TestSample1(t *testing.T) {
	s := `4 5
2 1
3 1
4 0
1 1
5 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 0
2 1
3 1
`
	expect := false
	runSample(t, s, expect)
}
