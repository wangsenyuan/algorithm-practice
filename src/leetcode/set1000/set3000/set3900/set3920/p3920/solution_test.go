package p3920

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxFixedPoints(nums)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", nums, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 2, 1}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 1, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 0, 1, 2}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{0, 3, 2}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{6, 4, 0, 4, 4, 5}
	expect := 2
	runSample(t, nums, expect)
}
