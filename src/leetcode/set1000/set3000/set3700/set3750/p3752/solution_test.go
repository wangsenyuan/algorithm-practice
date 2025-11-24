package p3752

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, target int, expect []int) {
	res := lexSmallestNegatedPerm(n, int64(target))
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	target := 0
	expect := []int{-3, 1, 2}
	runSample(t, n, target, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	target := 10000000000
	runSample(t, n, target, nil)
}