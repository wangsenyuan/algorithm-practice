package p3623

import "testing"

func runSample(t *testing.T, points [][]int, expect int) {
	res := countTrapezoids(points)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}
	expect := 3
	runSample(t, points, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}
	expect := 1
	runSample(t, points, expect)
}
