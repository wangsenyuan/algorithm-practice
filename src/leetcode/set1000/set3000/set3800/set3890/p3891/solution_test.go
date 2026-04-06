package p3891

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := minIncrease(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 1, 3}
	expect := int64(2)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{20, 7, 6, 4, 23, 17}
	expect := int64(2)
	runSample(t, nums, expect)
}
