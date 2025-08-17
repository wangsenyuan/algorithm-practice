package p3655

import "testing"

func runSample(t *testing.T, nums []int, queries [][]int, expect int) {
	res := xorAfterQueries(nums, queries)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 1}
	queries := [][]int{{0, 2, 1, 4}}
	expect := 4
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 1, 5, 4}
	queries := [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}}
	expect := 31
	runSample(t, nums, queries, expect)
}
