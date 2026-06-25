package main

import (
	"testing"
)

func runSample(t *testing.T, n int, k int, c int) {
	t.Helper()

	limit := n + 3
	var ops int

	S := make(map[int]bool)

	var initial int
	chooseInitial := func(a int) {
		if a < 0 || a >= 1<<n {
			t.Fatalf("initial element %d out of range", a)
		}
		initial = a
		clear(S)
		S[a] = true
	}

	it := Interactor{
		Insert: func(x int) int {
			ops++
			if ops > limit {
				t.Fatalf("too many operations: %d > %d", ops, limit)
			}
			if x < 0 || x >= 1<<n {
				t.Fatalf("insert argument %d out of range", x)
			}
			S[applyF(k, x, c)] = true
			return len(S)
		},
		Query: func(y int) int {
			ops++
			if ops > limit {
				t.Fatalf("too many operations: %d > %d", ops, limit)
			}
			if y < 0 || y >= 1<<n {
				t.Fatalf("query argument %d out of range", y)
			}
			var cnt int
			for z := range S {
				if z >= y {
					cnt++
				}
			}
			return cnt
		},
	}

	gotK, gotC := solve(n, chooseInitial, it)
	if gotK != k || gotC != c {
		t.Fatalf("Sample n=%d hidden (%d,%d), initial=%d, expect k=%d c=%d, but got k=%d c=%d",
			n, k, c, initial, k, c, gotK, gotC)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 1)
}

func TestAllSmallHiddenValues(t *testing.T) {
	for n := 2; n <= 6; n++ {
		for k := 1; k <= 3; k++ {
			for c := 1; c < 1<<n; c++ {
				runSample(t, n, k, c)
			}
		}
	}
}
