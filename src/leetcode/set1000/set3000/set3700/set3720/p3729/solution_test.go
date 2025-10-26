package p3729

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := numGoodSubarrays(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 3
	expect := int64(3)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2, 2}
	k := 6
	expect := int64(2)
	runSample(t, nums, k, expect)
}
