package p3544

import "testing"

func runSample(t *testing.T, edges [][]int, nums []int, k int, expect int64) {
	res := subtreeInversionSum(edges, nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{2, 6},
	}
	nums := []int{4, -8, -6, 3, 7, -2, 5}
	k := 2
	expect := int64(27)

	runSample(t, edges, nums, k, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	nums := []int{-1, 3, -2, 4, -5}
	k := 2
	expect := int64(9)

	runSample(t, edges, nums, k, expect)
}

func TestSample3(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{0, 2},
	}
	nums := []int{0, -1, -2}
	k := 3
	expect := int64(3)

	runSample(t, edges, nums, k, expect)
}

func TestSample4(t *testing.T) {
	edges := [][]int{
		{0, 1},
		{0, 2},
	}
	nums := []int{-5, 3, -5}
	k := 1
	expect := int64(13)

	runSample(t, edges, nums, k, expect)
}
