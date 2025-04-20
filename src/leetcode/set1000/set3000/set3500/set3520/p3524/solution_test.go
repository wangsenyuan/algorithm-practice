package p3524

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, queries [][]int, expect []int) {
	res := resultArray(nums, k, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	queries := [][]int{
		{2, 2, 0, 2},
		{3, 3, 3, 0},
		{0, 1, 0, 1},
	}
	expect := []int{2, 2, 2}
	runSample(t, nums, k, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4, 8, 16, 32}
	k := 4
	queries := [][]int{
		{0, 2, 0, 2},
		{0, 2, 0, 1},
	}
	expect := []int{1, 0}
	runSample(t, nums, k, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 9, 5, 7, 1}
	k := 2
	queries := [][]int{
		{5, 60, 4, 0},
		{1, 13, 5, 0},
		{0, 53, 4, 1},
		{3, 13, 0, 0},
	}
	// 53, 13, 9, 13, 7, 60
	expect := []int{1, 1, 1, 1}
	runSample(t, nums, k, queries, expect)
}
