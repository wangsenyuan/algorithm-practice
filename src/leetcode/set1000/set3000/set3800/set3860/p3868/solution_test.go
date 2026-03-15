package p3868

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := minCost(nums1, nums2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 1}
	nums2 := []int{2, 2}
	expect := 1
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 2, 2}
	nums2 := []int{1, 1, 1}
	expect := 1
	runSample(t, nums1, nums2, expect)
}
