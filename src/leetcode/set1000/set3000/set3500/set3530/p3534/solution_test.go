package p3534

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, maxDiff int, queries [][]int, expect []int) {
	res := pathExistenceQueries(len(nums), nums, maxDiff, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 8, 3, 4, 2}
	maxDiff := 3
	queries := [][]int{
		{0, 3}, {2, 4}, {0, 1},
	}
	expect := []int{1, 1, -1}
	runSample(t, nums, maxDiff, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{5, 3, 1, 9, 10}
	maxDiff := 2
	queries := [][]int{
		{0, 1}, {0, 2}, {2, 3}, {4, 3},
	}
	expect := []int{1, 2, -1, 1}
	runSample(t, nums, maxDiff, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 6, 1}
	maxDiff := 1
	queries := [][]int{
		{0, 0}, {0, 1}, {1, 2},
	}
	expect := []int{0, -1, -1}
	runSample(t, nums, maxDiff, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	maxDiff := 1
	queries := [][]int{
		{0, 5},
	}
	expect := []int{5}
	runSample(t, nums, maxDiff, queries, expect)
}
