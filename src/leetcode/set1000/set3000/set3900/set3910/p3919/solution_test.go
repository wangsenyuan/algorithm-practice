package p3919

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := minCost(nums, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, %v, expect %v, but got %v", nums, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{-5, -2, 3}
	queries := [][]int{{0, 2}, {2, 0}, {1, 2}}
	expect := []int{6, 2, 5}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 2, 3, 9}
	queries := [][]int{{3, 0}, {1, 2}, {2, 0}}
	expect := []int{4, 1, 3}
	runSample(t, nums, queries, expect)
}
