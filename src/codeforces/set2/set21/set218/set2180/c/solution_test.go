package main

import (
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	check := func(arr []int) int {
		var sum int
		var xor int
		for _, v := range arr {
			if v > n {
				t.Fatalf("Sample result %v, not correct", arr)
			}
			sum += v
			xor ^= v
		}
		if xor != n {
			t.Fatalf("Sample result %v, not correct", arr)
		}
		return sum
	}
	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 4, []int{1, 4, 5, 5})
	runSample(t, 8, 2, []int{0, 8})
}
