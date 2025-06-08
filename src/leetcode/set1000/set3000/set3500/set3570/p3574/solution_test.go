package p3574

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := maxGCDScore(nums, k)
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{2, 4}
	k := 1
	expect := 8
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 5, 7}
	k := 2
	expect := 14
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{5, 5, 5}
	k := 1
	expect := 15
	runSample(t, nums, k, expect)
}
