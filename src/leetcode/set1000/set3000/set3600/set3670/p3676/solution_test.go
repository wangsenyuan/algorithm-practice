package p3676

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := bowlSubarrays(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 5, 3, 1, 4}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{5, 1, 2, 3, 4}, 3)
}
