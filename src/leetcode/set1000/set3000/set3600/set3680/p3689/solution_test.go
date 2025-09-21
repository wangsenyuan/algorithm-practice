package p3689

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := minSplitMerge(nums1, nums2)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{3, 1, 2}
	nums2 := []int{1, 2, 3}
	expect := 1
	runSample(t, nums1, nums2, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 1, 1}
	expect := 3
	runSample(t, nums1, nums2, expect)
}
