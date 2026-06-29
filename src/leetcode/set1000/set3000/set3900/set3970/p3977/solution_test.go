package p3977

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, power int, cost []int, source int, target int, expect []int64) {
	res := minTimeMaxPower(n, edges, power, cost, source, target)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1, 1}, {1, 4, 1}, {0, 2, 1}, {2, 3, 1}, {3, 4, 1},
	}
	power := 4
	cost := []int{2, 3, 1, 1, 1}
	source := 0
	target := 4
	expect := []int64{3, 0}
	runSample(t, n, edges, power, cost, source, target, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1, 2}, {1, 2, 2}, {2, 0, 2},
	}
	power := 3
	cost := []int{1, 1, 1}
	source := 1
	target := 1
	expect := []int64{0, 3}
	runSample(t, n, edges, power, cost, source, target, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1, 3}, {2, 3, 4},
	}
	power := 3
	cost := []int{1, 1, 1, 1}
	source := 0
	target := 3
	expect := []int64{-1, -1}
	runSample(t, n, edges, power, cost, source, target, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	edges := [][]int{
		{0, 1, 9},
	}
	power := 3
	cost := []int{2, 1}
	source := 1
	target := 0
	expect := []int64{-1, -1}
	runSample(t, n, edges, power, cost, source, target, expect)
}

func TestSample5(t *testing.T) {
	n := 10
	edges := [][]int{
		{8, 1, 27}, {1, 0, 10}, {9, 2, 22}, {0, 8, 28}, {6, 7, 97},
		{9, 3, 36}, {3, 7, 33}, {8, 7, 35}, {5, 3, 79}, {9, 7, 62},
		{1, 2, 17}, {6, 4, 48}, {4, 1, 32}, {1, 2, 73}, {2, 4, 12},
		{6, 1, 41}, {6, 0, 92}, {0, 1, 48}, {5, 2, 39}, {5, 9, 47},
		{2, 6, 26}, {6, 7, 100}, {8, 3, 63}, {9, 5, 43}, {1, 7, 34},
		{4, 0, 61}, {2, 9, 36}, {4, 4, 57}, {6, 1, 49}, {2, 6, 79},
		{1, 2, 100}, {1, 1, 18}, {9, 4, 67}, {9, 5, 54}, {4, 1, 98},
		{7, 5, 23}, {9, 5, 12}, {5, 5, 13}, {6, 9, 48}, {8, 7, 29},
		{9, 4, 61}, {7, 4, 80}, {4, 2, 37}, {7, 8, 70}, {7, 2, 95},
		{0, 1, 11}, {0, 9, 74}, {3, 0, 16}, {4, 3, 22}, {1, 0, 97},
		{5, 5, 98}, {6, 9, 24}, {6, 4, 79}, {2, 2, 72}, {0, 8, 43},
		{2, 5, 57}, {8, 7, 69}, {3, 1, 29}, {8, 2, 97}, {0, 9, 97},
		{3, 7, 38}, {6, 8, 50}, {7, 4, 75}, {6, 9, 13}, {1, 1, 51},
		{1, 6, 29}, {0, 9, 67}, {3, 2, 76}, {1, 9, 32}, {2, 3, 85},
		{6, 4, 45}, {8, 6, 44}, {9, 4, 70},
	}
	power := 58
	cost := []int{3, 6, 66, 2, 9, 6, 20, 6, 15, 9}
	source := 8
	target := 0
	expect := []int64{37, 37}
	runSample(t, n, edges, power, cost, source, target, expect)
}
