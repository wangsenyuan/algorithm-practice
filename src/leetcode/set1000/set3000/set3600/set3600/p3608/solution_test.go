package p3608

import "testing"

func runSample(t *testing.T, n int, edges [][]int, k int, expect int) {
	res := minTime(n, edges, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{{0, 1, 3}}
	k := 2
	expect := 3
	runSample(t, n, edges, k, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 2}, {1, 2, 4}}
	k := 3
	expect := 4
	runSample(t, n, edges, k, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	edges := [][]int{{0, 2, 5}}
	k := 2
	expect := 0
	runSample(t, n, edges, k, expect)
}
