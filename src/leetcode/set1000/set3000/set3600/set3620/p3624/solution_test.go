package p3624

import "testing"

func runSample(t *testing.T, nums []int64, queries [][]int64, expect []int) {
	res := popcountDepth(nums, queries)

	for i := range expect {
		if res[i] != expect[i] {
			t.Errorf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	nums := []int64{2, 4}
	queries := [][]int64{{1, 0, 1, 1}, {2, 1, 1}, {1, 0, 1, 0}}
	expect := []int{2, 1}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int64{3, 5, 6}
	queries := [][]int64{{1, 0, 2, 2}, {2, 1, 4}, {1, 1, 2, 1}, {1, 0, 1, 0}}
	expect := []int{3, 1, 0}
	runSample(t, nums, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int64{1, 2}
	queries := [][]int64{{1, 0, 1, 1}, {2, 0, 3}, {1, 0, 0, 1}, {1, 0, 0, 2}}
	expect := []int{1, 0, 1}
	runSample(t, nums, queries, expect)
}
