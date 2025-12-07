package p3768

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := minInversionCount(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 2, 5, 4}
	k := 3
	expect := int64(0)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 3, 2, 1}
	k := 4
	expect := int64(6)
	runSample(t, nums, k, expect)
}
