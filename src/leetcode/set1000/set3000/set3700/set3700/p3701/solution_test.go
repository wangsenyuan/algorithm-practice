package p3701

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestSubsequence(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 2
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 4}
	expect := 3
	runSample(t, nums, expect)
}
