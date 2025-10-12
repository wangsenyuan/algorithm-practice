package p3715

import "testing"

func runSample(t *testing.T, n int, edges [][]int, nums []int, expect int64) {
	res := sumOfAncestors(n, edges, nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	nums := []int{2, 8, 2}
	var expect int64 = 3
	runSample(t, n, edges, nums, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
	}
	nums := []int{1, 2, 4}
	var expect int64 = 1
	runSample(t, n, edges, nums, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
	}
	nums := []int{1, 2, 9, 4}
	var expect int64 = 2
	runSample(t, n, edges, nums, expect)
}
