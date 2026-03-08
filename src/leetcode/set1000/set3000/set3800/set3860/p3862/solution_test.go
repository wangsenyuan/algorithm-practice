package p3862

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := smallestBalancedIndex(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 2}
	expect := 1
	runSample(t, nums, expect)
}
