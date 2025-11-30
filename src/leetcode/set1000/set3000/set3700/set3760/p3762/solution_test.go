package p3762

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, k int, queries [][]int, expect []int64) {
	res := minOperations(nums, k, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 4, 7}
	k := 3
	queries := [][]int{{0, 1}, {0, 2}}
	expect := []int64{1, 2}
	runSample(t, nums, k, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 4}
	k := 2
	queries := [][]int{{0, 2}, {0, 0}, {1, 2}}
	expect := []int64{-1, 0, 1}
	runSample(t, nums, k, queries, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{95808, 64116, 28845}
	k := 2
	queries := [][]int{{1, 2}}
	expect := []int64{-1}
	runSample(t, nums, k, queries, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{82, 69, 12, 50, 89, 21, 34, 96, 61, 59}
	k := 10
	queries := [][]int{{0, 3}, {3, 3}, {1, 9}, {7, 8}, {3, 6}, {6, 6}, {5, 5}, {0, 2}, {5, 7}, {4, 8}}
	expect := []int64{-1, 0, -1, -1, -1, 0, 0, -1, -1, -1}
	runSample(t, nums, k, queries, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{72963, 3078, 113817, 49068, 132672, 61722, 60654, 113877, 31095, 1203, 97263, 42228, 105450, 139155, 4161, 41436, 8853, 112602, 135624, 81423, 41676, 37491, 102135, 79890, 34620, 240, 125955, 85539, 40782}
	k := 3
	queries := [][]int{{21, 26}, {15, 20}, {14, 15}, {17, 18}, {18, 27}, {16, 20}, {4, 22}, {13, 14}, {26, 28}, {23, 25}, {19, 21}, {23, 28}, {21, 28}, {17, 24}, {10, 16}, {8, 22}, {6, 7}, {11, 13}, {9, 14}, {23, 27}, {3, 27}}
	expect := []int64{78543, 79228, 12425, 7674, 112253, 65899, 250468, 44998, 28391, 26550, 14644, 71914, 93462, 79369, 95806, 202579, 17741, 32309, 98092, 58878, 319620}
	runSample(t, nums, k, queries, expect)
}
