package main

import "testing"

func runSample(t *testing.T, nums []int, k int, limit int, expect int) {
	res := maxProduct(nums, k, limit)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3}
	k := 2
	limit := 10
	expect := 6
	runSample(t, nums, k, limit, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{0, 2, 3}
	k := -5
	limit := 12
	expect := -1
	runSample(t, nums, k, limit, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 3, 3}
	k := 0
	limit := 9
	expect := 9
	runSample(t, nums, k, limit, expect)
}
