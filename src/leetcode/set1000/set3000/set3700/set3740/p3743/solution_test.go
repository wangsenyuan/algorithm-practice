package p3743

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int64) {
	res := maximumScore(nums, k)
	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	k := 2
	expect := int64(3)
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	k := 1
	expect := int64(2)
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	k := 4
	expect := int64(3)
	runSample(t, nums, k, expect)
}
