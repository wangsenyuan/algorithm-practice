package p3882

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := minCost(grid)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{{1, 2}, {3, 4}}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{{6, 7}, {5, 8}}
	expect := 9
	runSample(t, grid, expect)
}
