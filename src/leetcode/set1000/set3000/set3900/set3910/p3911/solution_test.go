package p3911

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := kthRemainingInteger(nums, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 7}
	queries := [][]int{{0, 2, 1}, {1, 1, 2}, {0, 0, 3}}
	expect := []int{2, 6, 6}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 5, 8}
	queries := [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}
	expect := []int{6, 2, 12}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 6}
	queries := [][]int{{0, 1, 1}, {1, 1, 3}}
	expect := []int{2, 8}
	runSample(t, nums, queries, expect)
}
