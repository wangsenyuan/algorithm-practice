package p3627

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maximumMedianSum(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 1, 3, 2, 1, 3}, 5)
}
