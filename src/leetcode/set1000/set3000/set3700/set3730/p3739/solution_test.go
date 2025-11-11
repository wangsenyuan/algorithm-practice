package p3739

import "testing"

func runSample(t *testing.T, nums []int, target int, expect int64) {
	res := countMajoritySubarrays(nums, target)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	target := 2
	expect := int64(5)
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	target := 1
	expect := int64(10)
	runSample(t, nums, target, expect)
}
