package main

import "testing"

func runSample(t *testing.T, nums []int, expect int) {
	res := solve(nums)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 1, 1, 4}
	runSample(t, nums, 5)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 2, 1, 4}
	expect := 4
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{4, 1, 2}
	expect := 0
	runSample(t, nums, expect)
}
