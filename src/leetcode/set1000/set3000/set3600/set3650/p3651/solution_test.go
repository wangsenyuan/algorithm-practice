package p3651

import "testing"

func runSample(t *testing.T, grid [][]int, k int, expect int) {
	res := minCost(grid, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}}
	k := 2
	expect := 7
	runSample(t, grid, k, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{1, 2}, {2, 3}, {3, 4}}
	k := 1
	expect := 9
	runSample(t, grid, k, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{{0, 12, 1, 14, 3}, {8, 18, 18, 4, 8}, {15, 17, 13, 16, 6}, {11, 16, 23, 15, 16}, {21, 17, 25, 27, 24}, {26, 27, 21, 16, 23}, {27, 20, 31, 30, 32}}
	k := 2
	expect := 78
	runSample(t, grid, k, expect)
}
