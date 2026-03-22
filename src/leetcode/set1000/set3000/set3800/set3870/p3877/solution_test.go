package p3877

import "testing"

func runSample(t *testing.T, nums []int, target int, expect int) {
	res := minRemovals(nums, target)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 2
	expect := 1
	runSample(t, nums, target, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 4}
	target := 1
	expect := -1
	runSample(t, nums, target, expect)
}
