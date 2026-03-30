package p3887

import "testing"

func runSample(t *testing.T, n int, edges [][]int, expect int) {
	res := numberOfEdgesAdded(n, edges)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, edges := 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 1}}
	expect := 2
	runSample(t, n, edges, expect)
}

func TestSample2(t *testing.T) {
	n, edges := 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 0}}
	expect := 3
	runSample(t, n, edges, expect)
}

func TestSample3(t *testing.T) {
	n, edges := 3, [][]int{{0, 1, 0}}
	expect := 1
	runSample(t, n, edges, expect)
}
