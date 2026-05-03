package lcr039

import "testing"

func runSample(t *testing.T, heights []int, expect int) {
	res := largestRectangleArea(heights)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", heights, expect, res)
	}
}

func TestSample1(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	expect := 10
	runSample(t, heights, expect)
}

func TestSample2(t *testing.T) {
	heights := []int{2, 4}
	expect := 4
	runSample(t, heights, expect)
}

func TestSample3(t *testing.T) {
	heights := []int{4, 2, 0, 3, 2, 4, 3, 4}
	expect := 10
	runSample(t, heights, expect)
}
