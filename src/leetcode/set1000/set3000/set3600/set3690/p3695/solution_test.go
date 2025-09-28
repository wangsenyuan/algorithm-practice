package p3695

import "testing"

func runSample(t *testing.T, nums []int, swaps [][]int, expect int64) {
	res := maxAlternatingSum(nums, swaps)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	swaps := [][]int{{0, 2}, {1, 2}}
	expect := int64(4)
	runSample(t, nums, swaps, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1000000000, 1, 1000000000, 1, 1000000000}
	expect := int64(-2999999997)
	runSample(t, nums, nil, expect)
}
