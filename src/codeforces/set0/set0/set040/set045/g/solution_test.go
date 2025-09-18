package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if expect < 0 {
		if len(res) != 0 {
			t.Fatalf("Sample expect %d, but got %v", expect, res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	nums := slices.Clone(res)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	if len(nums) != expect {
		t.Fatalf("Sample expect %d colors, but got %v", expect, nums)
	}
	sum := make([]int, expect+1)
	for i, v := range res {
		sum[v] += i + 1
	}
	for i := 1; i <= expect; i++ {
		if !checkPrimeFast(sum[i]) {
			t.Fatalf("Sample result %v, not correct, its sum (numbers) of color %d, is not a prime, which is %d", res, i, sum[i])
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 2)
}