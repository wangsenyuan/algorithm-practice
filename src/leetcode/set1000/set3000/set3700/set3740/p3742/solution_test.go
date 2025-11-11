package p3742

import "testing"

func runSample(t *testing.T, grid [][]int, k int, expect int) {
	res := maxPathScore(grid, k)
	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{0, 1}, {2, 0}}
	k := 1
	expect := 2
	runSample(t, grid, k, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{0, 1}, {1, 2}}
	k := 1
	expect := -1
	runSample(t, grid, k, expect)
}
