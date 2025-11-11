package p3738

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestSubarray(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, -4, -2}
	expect := 3
	runSample(t, nums, expect)
}
