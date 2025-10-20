package p3721

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := longestBalanced(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 5, 4, 3}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 2, 2, 5, 4}
	expect := 5
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 2}
	expect := 3
	runSample(t, nums, expect)
}
