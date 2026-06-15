package p3962

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxSum(nums, k)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, -1, 0, 2}
	k := 1
	expect := 3
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 3, 2, 4}
	k := 2
	expect := 13
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{-1, -2}
	k := 1
	expect := -1
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{-49, 37, 42, -36, 37}
	k := 2
	expect := 116
	runSample(t, nums, k, expect)
}
