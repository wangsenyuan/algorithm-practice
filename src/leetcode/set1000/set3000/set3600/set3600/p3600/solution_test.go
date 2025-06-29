package p3600

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, expect int) {
	res := maxStability(n, edges, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 2, 1}, {1, 2, 3, 0}}
	k := 1
	expect := 2
	runSample(t, n, edges, k, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 4, 0}, {1, 2, 3, 0}, {0, 2, 1, 0}}
	k := 2
	expect := 6
	runSample(t, n, edges, k, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 1, 1}, {1, 2, 1, 1}, {2, 0, 1, 1}}
	k := 0
	expect := -1
	runSample(t, n, edges, k, expect)
}
