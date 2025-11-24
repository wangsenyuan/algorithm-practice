package p3755

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := maxBalancedSubarray(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 3, 2, 0}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 8, 5, 4, 14, 9, 15}
	expect := 8
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 4, 4, 4, 2, 1, 4, 1, 1, 2, 2, 4, 4, 3, 1, 4, 1, 1, 2, 1, 2, 2, 0, 1, 4, 1, 2, 0, 1, 1, 1, 2, 0, 0, 0, 2, 2, 4, 3, 2, 3, 1, 0, 2, 3, 2, 4, 1, 4, 0, 2, 0, 4, 4, 3, 3, 2}
	expect := 20
	runSample(t, nums, expect)
}
