package p3914

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := minOperations(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 3, 2, 1}
	expect := int64(2)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 1, 2, 3}
	expect := int64(4)
	runSample(t, nums, expect)
}
