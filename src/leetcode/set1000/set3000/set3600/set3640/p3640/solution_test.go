package p3640

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := maxSumTrionic(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, -2, -1, -3, 0, 2, -1}
	expect := int64(-4)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 4, 2, 7}
	expect := int64(14)
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{467, 941, -696, -288}
	expect := int64(424)
	runSample(t, nums, expect)
}
