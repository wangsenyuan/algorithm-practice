package p3522

import (
	"testing"
)

func runSample(t *testing.T, nums []int, expect int) {
	res := maximumPossibleSize(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{4, 2, 5, 3, 5}
	expect := 3
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := 3
	runSample(t, nums, expect)
}
