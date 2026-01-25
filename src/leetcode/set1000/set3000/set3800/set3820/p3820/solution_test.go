package p3820

import "testing"

func runSample(t *testing.T, n int, edges [][]int, x int, y int, z int, expect int) {
	res := specialNodes(n, edges, x, y, z)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}}
	x := 1
	y := 2
	z := 3
	expect := 3
	runSample(t, n, edges, x, y, z, expect)
}
