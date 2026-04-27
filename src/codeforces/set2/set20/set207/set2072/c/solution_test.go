package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, x int, expect []int) {
	check := func(res []int) int {
		slices.Sort(res)
		var mex int
		var or int
		for _, v := range res {
			or |= v
			if v == mex {
				mex++
			}
		}
		if or != x {
			t.Fatalf("Sample result %v, not correct", res)
		}
		return mex
	}

	res := solve(n, x)

	u := check(expect)
	v := check(res)

	if u != v {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 1, 69
	expect := []int{69}
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n, x := 7, 7
	expect := []int{6, 0, 3, 4, 1, 2, 5}
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n, x := 5, 7
	expect := []int{4, 1, 3, 0, 2}
	runSample(t, n, x, expect)
}

func TestSample4(t *testing.T) {
	n, x := 7, 3
	expect := []int{0, 1, 2, 3, 2, 1, 0}
	runSample(t, n, x, expect)
}

func TestSample5(t *testing.T) {
	n, x := 8, 7
	expect := []int{7, 0, 6, 1, 5, 2, 4, 3}
	runSample(t, n, x, expect)
}

func TestSample6(t *testing.T) {
	n, x := 3, 52
	expect := []int{0, 52, 0}
	runSample(t, n, x, expect)
}

func TestSample7(t *testing.T) {
	n, x := 9, 11
	expect := []int{0, 1, 8, 3, 0, 9, 11, 2, 10}
	runSample(t, n, x, expect)
}

func TestSample8(t *testing.T) {
	n, x := 6, 15
	expect := []int{4, 0, 3, 8, 1, 2}
	runSample(t, n, x, expect)
}

func TestSample9(t *testing.T) {
	n, x := 2, 3
	expect := []int{0, 3}
	runSample(t, n, x, expect)
}

