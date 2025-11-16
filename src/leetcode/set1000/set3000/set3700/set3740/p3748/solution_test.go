package p3748

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int64) {
	res := countStableSubarrays(nums, queries)

	if !slices.Equal(res, expect) {
		t.Errorf("Sample failed, expect %v, got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 1, 2}
	queries := [][]int{{0, 1}, {1, 2}, {0, 2}}
	expect := []int64{2, 3, 4}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 2}
	queries := [][]int{{0, 1}, {0, 0}}
	expect := []int64{3, 1}
	runSample(t, nums, queries, expect)
}
