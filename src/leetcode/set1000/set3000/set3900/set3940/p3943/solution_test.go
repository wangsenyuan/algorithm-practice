package p3943

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums1 []int, nums2 []int, queries [][]int, expect []int) {
	res := numberOfPairs(nums1, nums2, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample %v %v %v, expect %v, but got %v", nums1, nums2, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	queries := [][]int{{2, 5}, {1, 0, 0, 2}, {2, 5}}
	expect := []int{2, 1}
	runSample(t, nums1, nums2, queries, expect)
}

func TestSample2(t *testing.T) {
	nums1 := []int{1, 1}
	nums2 := []int{2, 2, 3}
	queries := [][]int{{2, 4}, {1, 0, 1, 1}, {2, 4}}
	expect := []int{2, 6}
	runSample(t, nums1, nums2, queries, expect)
}
