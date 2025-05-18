package p3552

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := minMoves(grid)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"A..", ".A.", "...",
	}
	expect := 2
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		".#...", ".#.#.", ".#.#.", "...#.",
	}
	expect := 13
	runSample(t, grid, expect)
}
