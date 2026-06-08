package p3956

import "testing"

func runSample(t *testing.T, nums []int, m int, l int, r int, expect int64) {
	res := maximumSum(nums, m, l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 1, -5, 2}
	runSample(t, nums, 2, 1, 3, 7)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 3, 4}
	runSample(t, nums, 2, 1, 2, 8)
}

func TestSample3(t *testing.T) {
	nums := []int{-1, 7, -4}
	runSample(t, nums, 1, 2, 3, 6)
}

func TestSample4(t *testing.T) {
	nums := []int{-3, -4, -1}
	runSample(t, nums, 2, 1, 2, -1)
}

func TestLimitNumberOfSubarrays(t *testing.T) {
	nums := []int{10, -100, 10, -100, 10}
	runSample(t, nums, 2, 1, 1, 20)
}

func TestSingleNegativeElement(t *testing.T) {
	nums := []int{-5}
	runSample(t, nums, 1, 1, 1, -5)
}
