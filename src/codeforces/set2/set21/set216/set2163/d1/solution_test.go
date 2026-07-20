package main

import (
	"testing"
)

func mex(vals []int) int {
	seen := make([]bool, len(vals)+1)
	for _, v := range vals {
		if v < len(seen) {
			seen[v] = true
		}
	}
	for i := range seen {
		if !seen[i] {
			return i
		}
	}
	return len(seen)
}

func runSample(t *testing.T, p []int, ranges [][]int, expect int) {
	t.Helper()
	n := len(p)
	limit := 300
	if half := (n+1)/2 + 2; half > limit {
		limit = half
	}
	var cnt int
	ask := func(l, r int) int {
		cnt++
		if cnt > limit {
			t.Fatalf("too many queries: %d > %d", cnt, limit)
		}
		if l < 1 || r > n || l > r {
			t.Fatalf("invalid query [%d, %d]", l, r)
		}
		return mex(p[l-1 : r])
	}
	got := solve(n, ranges, ask)
	if got != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, got)
	}
}

func expectedMaxMex(p []int, ranges [][]int) int {
	ans := 0
	for _, rg := range ranges {
		m := mex(p[rg[0]-1 : rg[1]])
		if m > ans {
			ans = m
		}
	}
	return ans
}

func TestSample1(t *testing.T) {
	// Hidden p from the statement sketch for first test case shape.
	p := []int{1, 0, 3, 2}
	ranges := [][]int{{1, 2}, {2, 4}, {1, 3}}
	runSample(t, p, ranges, expectedMaxMex(p, ranges))
}

func TestSample2(t *testing.T) {
	p := []int{3, 2, 0, 1, 5, 4}
	ranges := [][]int{{1, 2}, {2, 4}, {3, 3}, {4, 6}, {5, 5}, {6, 6}}
	runSample(t, p, ranges, expectedMaxMex(p, ranges))
}

func TestSample3(t *testing.T) {
	p := []int{0, 1, 2, 3}
	ranges := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	runSample(t, p, ranges, expectedMaxMex(p, ranges))
}

func TestTernarySearchPlateauBeforeMaximum(t *testing.T) {
	p := []int{4, 5, 6, 0, 2, 3, 1}
	ranges := [][]int{{1, 4}, {2, 5}, {3, 6}, {4, 7}}
	runSample(t, p, ranges, 4)
}
