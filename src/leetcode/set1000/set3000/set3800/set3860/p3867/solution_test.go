package p3867

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := gcdSum(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 6, 2, 8}
	expect := int64(5)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 6, 4}
	expect := int64(2)
	runSample(t, nums, expect)
}
