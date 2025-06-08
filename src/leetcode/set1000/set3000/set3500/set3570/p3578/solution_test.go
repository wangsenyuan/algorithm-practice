package p3578

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := countPartitions(nums, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{9, 4, 1, 3, 7}, 4, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{3, 3, 4}, 0, 2)
}
