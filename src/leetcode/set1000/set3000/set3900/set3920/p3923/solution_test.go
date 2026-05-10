package p3923

import "testing"

func runSample(t *testing.T, points [][]int, target []int, expect int) {
	res := minGenerations(points, target)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{0, 0, 0}, {6, 6, 6}}
	target := []int{3, 3, 3}
	expect := 1
	runSample(t, points, target, expect)
}
