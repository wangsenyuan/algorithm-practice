package p3551

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := minSwaps(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{37, 100}
	expect := 1
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{22, 14, 33, 7}
	expect := 0
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{18, 43, 34, 16}
	expect := 2
	runSample(t, nums, expect)
}
