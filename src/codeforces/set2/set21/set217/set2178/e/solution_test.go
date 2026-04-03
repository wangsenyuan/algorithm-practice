package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, a []int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	ask := func(l int, r int) int {
		return sum[r] - sum[l-1]
	}

	res := solve(n, ask)

	expect := slices.Max(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1, 2, 2, 2, 1, 1, 2, 2}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []int{4, 4, 4, 4, 4, 4, 4, 4}
	runSample(t, a)
}

// Third test case in the statement example: n = 4, every a_i = 4 (queries ? 1 1 … ? 4 4, answer ! 4).
func TestSample3(t *testing.T) {
	a := []int{4, 4, 4, 4}
	runSample(t, a)
}

// Fourth test case in the statement example: answer ! 1073741824 (2^30).
func TestSample4(t *testing.T) {
	const (
		p28 = 1 << 28
		p29 = 1 << 29
		p30 = 1 << 30
	)
	a := []int{p29, p29, p30, p29, p28, p28, p29, p29}
	runSample(t, a)
}
