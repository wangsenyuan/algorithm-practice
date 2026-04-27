package p3915

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := maxAlternatingSum(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{5, 4, 2}
	k := 2
	expect := int64(7)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 5, 4, 2, 4}
	k := 1
	expect := int64(14)
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{5}
	k := 1
	expect := int64(5)
	runSample(t, nums, k, expect)
}
