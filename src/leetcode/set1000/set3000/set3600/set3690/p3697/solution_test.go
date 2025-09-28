package p3697

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := splitArray(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 3, 2}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 2, 4, 3}, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{3, 1, 2}, -1)
}
