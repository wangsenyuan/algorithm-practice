package p3636

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := subarrayMajority(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 2, 1, 1}
	queries := [][]int{{0, 5, 4}, {0, 3, 3}, {2, 3, 2}}
	expect := []int{1, -1, 2}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 3, 2, 3, 2, 3}
	queries := [][]int{{0, 6, 4}, {1, 5, 2}, {2, 4, 1}, {3, 3, 1}}
	expect := []int{3, 2, 3, 2}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{15}
	queries := [][]int{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}, {0, 0, 1}}
	expect := []int{15, 15, 15, 15, 15, 15, 15, 15, 15}
	runSample(t, nums, queries, expect)
}
