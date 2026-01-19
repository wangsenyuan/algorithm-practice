package p3812

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, start string, target string, expect []int) {
	res := minimumFlips(n, edges, start, target)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	start := "010"
	target := "100"
	expect := []int{0}
	runSample(t, n, edges, start, target, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {3, 5}, {1, 6}}
	start := "0011000"
	target := "0010001"
	expect := []int{1, 2, 5}
	runSample(t, n, edges, start, target, expect)
}
