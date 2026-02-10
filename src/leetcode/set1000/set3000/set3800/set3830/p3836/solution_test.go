package p3836

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, k int, expect int64) {
	res := maxScore(nums1, nums2, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 3, 2}
	nums2 := []int{4, 5, 1}
	k := 2
	expect := int64(22)
	runSample(t, nums1, nums2, k, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{-2, 0, 5}
	nums2 := []int{-3, 4, -1, 2}
	k := 2
	expect := int64(26)
	runSample(t, nums1, nums2, k, expect)
}

func TestSample3(t *testing.T) {
	nums1 := []int{-3, -2}
	nums2 := []int{1, 2}
	k := 2
	expect := int64(-7)
	runSample(t, nums1, nums2, k, expect)
}
