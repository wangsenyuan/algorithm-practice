package p3686

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := countStableSubsequences(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 5}
	expect := 6
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{2, 3, 4, 2}
	expect := 14
	runSample(t, nums, expect)
}
