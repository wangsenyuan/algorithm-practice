package p3976

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxSubarraySum(nums, k)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -2, 3, 4, -5}
	k := 2
	expect := 14
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{-5, -4, -3}
	k := 2
	expect := -1
	runSample(t, nums, k, expect)
}
