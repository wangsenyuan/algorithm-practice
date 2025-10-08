package p3700

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := alternatingSum(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 5, 7}
	expect := -4
	runSample(t, nums, expect)
}
