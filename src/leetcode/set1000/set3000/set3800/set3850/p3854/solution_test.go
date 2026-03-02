package p3854

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := makeParityAlternating(nums)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{-2, -3, 1, 4}
	expect := []int{2, 6}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 2, -2}
	expect := []int{1, 3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{7, 7}
	expect := []int{1, 1}
	runSample(t, nums, expect)
}
