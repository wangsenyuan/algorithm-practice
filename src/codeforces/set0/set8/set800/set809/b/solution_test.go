package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, a []int) {
	var times int

	find := func(x int) int {
		if x == 0 || x > n {
			t.Fatalf("find(%d) is out of range", x)
		}
		i := sort.SearchInts(a, x)
		dx := 1 << 60
		if i < len(a) {
			dx = a[i] - x
		}
		if i > 0 {
			dx = min(dx, x-a[i-1])
		}
		return dx
	}

	ask := func(x int, y int) string {
		times++
		if times > 60 {
			t.Fatalf("Sample asked too much")
		}

		dx := find(x)
		dy := find(y)
		if dx <= dy {
			return "TAK"
		}
		return "NIE"
	}

	res := solve(n, len(a), ask)
	x, y := res[0], res[1]
	if x == y {
		t.Fatalf("Sample result %v, x == y", res)
	}
	if find(x) != 0 || find(y) != 0 {
		t.Fatalf("Sample result %v, x or y is not the answer", res)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{2, 3}
	runSample(t, n, a)
}

func TestSample2(t *testing.T) {
	n := 100
	a := []int{3, 66}
	runSample(t, n, a)
}

func TestSample3(t *testing.T) {
	n := 3
	a := []int{1, 2}
	runSample(t, n, a)
}

