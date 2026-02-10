package p3835

import "testing"

func runSample(t *testing.T, nums []int, k int64, expect int64) {
	res := countSubarrays(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 2}
	k := int64(4)
	expect := int64(5)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 5, 5, 5}
	k := int64(0)
	expect := int64(10)
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3}
	k := int64(0)
	expect := int64(3)
	runSample(t, nums, k, expect)
}
