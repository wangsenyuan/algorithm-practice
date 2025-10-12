package p3710

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := maxPartitionFactor(points)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	expect := 4
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{0, 0}, {0, 1}, {10, 0}}
	expect := 11
	runSample(t, points, expect)
}
