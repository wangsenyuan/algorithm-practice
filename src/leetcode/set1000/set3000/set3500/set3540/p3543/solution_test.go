package p3543

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, w int, expect int) {
	res := maxWeight(n, edges, k, w)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 1},
		{1, 2, 2},
	}
	k := 2
	w := 4
	expect := 3

	runSample(t, n, edges, k, w, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 2},
		{0, 2, 3},
	}
	k := 1
	w := 3
	expect := 2

	runSample(t, n, edges, k, w, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 6},
		{1, 2, 8},
	}
	k := 1
	w := 6
	expect := -1

	runSample(t, n, edges, k, w, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 9},
		{0, 2, 4},
		{0, 1, 7},
	}
	k := 2
	w := 331
	expect := 16

	runSample(t, n, edges, k, w, expect)
}
