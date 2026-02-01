package p3830

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestAlternating(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3, 2}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 1, 2, 3, 2, 1}
	expect := 4
	runSample(t, nums, expect)
}
