package p3569

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries [][]int, expect []int) {
	res := maximumCount(nums, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 1, 3, 1, 2}
	queries := [][]int{{1, 2}, {3, 3}}
	expect := []int{3, 4}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 1, 4}
	queries := [][]int{{0, 1}}
	expect := []int{0}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 2}
	queries := [][]int{{0, 2}, {0, 2}}
	expect := []int{2, 2}
	runSample(t, nums, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{3, 18, 3}
	queries := [][]int{{1, 57}, {1, 3}, {0, 91}}
	expect := []int{2, 2, 2}
	runSample(t, nums, queries, expect)
}
