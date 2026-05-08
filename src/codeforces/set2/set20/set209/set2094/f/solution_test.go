package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int) {
	cnt := make([]int, k)

	res := solve(n, m, k)

	for i := range n {
		for j := range m {
			if i > 0 && res[i-1][j] == res[i][j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if j > 0 && res[i][j-1] == res[i][j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			cnt[res[i][j]-1]++
		}
	}
	x := slices.Min(cnt)
	y := slices.Max(cnt)
	if x != y {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 4, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 3, 2)
}

func TestSample6(t *testing.T) {
	/**
	1 2
	3 2
	3 1
	*/
	runSample(t, 3, 2, 3)
}