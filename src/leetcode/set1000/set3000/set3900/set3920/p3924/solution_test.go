package p3924

import "testing"

func runSample(t *testing.T, n int, edges [][]int, source int, target int, k int, expect int) {
	res := minimumThreshold(n, edges, source, target, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1, 5}, {1, 2, 3}, {3, 4, 4}, {4, 5, 1}, {1, 4, 2}}
	source := 0
	target := 3
	k := 1
	expect := 4
	runSample(t, n, edges, source, target, k, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	edges := [][]int{{0, 1, 3}, {1, 2, 4}, {3, 4, 5}, {4, 5, 6}}
	source := 0
	target := 4
	k := 1
	expect := -1
	runSample(t, n, edges, source, target, k, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1, 2}, {1, 2, 2}, {2, 3, 2}, {3, 0, 2}}
	source := 0
	target := 0
	k := 0
	expect := 0
	runSample(t, n, edges, source, target, k, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	edges := [][]int{{2, 3, 18}, {1, 2, 96}}
	source := 1
	target := 2
	k := 2
	expect := 0
	runSample(t, n, edges, source, target, k, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	edges := [][]int{{1, 2, 77}, {0, 2, 36}}
	source := 1
	target := 1
	k := 1
	expect := 0
	runSample(t, n, edges, source, target, k, expect)
}
