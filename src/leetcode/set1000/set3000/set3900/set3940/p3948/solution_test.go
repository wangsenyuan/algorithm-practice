package p3948

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, nums []int, expect []int) {
	res := maximumMEX(nums)

	if !slices.Equal(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{0, 1, 0}
	expect := []int{2, 1}
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 0, 2}
	expect := []int{3}
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{3, 1}
	expect := []int{0, 0}
	runSample(t, nums, expect)
}
