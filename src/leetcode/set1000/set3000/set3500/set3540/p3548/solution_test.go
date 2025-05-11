package p3548

import "testing"

func runSample(t *testing.T, grid [][]int, expect bool) {
	res := canPartitionGrid(grid)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]int{
		{1, 4},
		{2, 3},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := true
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := [][]int{
		{1, 2, 4},
		{2, 3, 5},
	}
	expect := false
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := [][]int{
		{4, 1, 8},
		{3, 2, 6},
	}
	expect := false
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := [][]int{
		{4, 1, 2, 2, 2},
	}
	expect := false
	runSample(t, grid, expect)
}

func TestSample6(t *testing.T) {
	grid := [][]int{
		{3, 1, 1},
	}
	expect := true
	runSample(t, grid, expect)
}
