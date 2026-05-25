package p3938

import "testing"

func runSample(t *testing.T, grid [][]int, expect int) {
	res := maxScore(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{4, -2, -3},
		{-1, -3, -1},
		{-4, 2, -1},
	}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2, 0, -3}, {1, -2, 1, 0}, {-4, 2, -1, 3}, {3, -3, 3, -2}, {-1, -5, 0, 1},
	}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{-5, 5, -5},
		{-5, -5, -5},
	}
	expect := 0
	runSample(t, grid, expect)
}
