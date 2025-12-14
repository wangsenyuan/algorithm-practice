package p3777

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := minDeletions(s, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ABA"
	queries := [][]int{{2, 1, 2}, {1, 1}, {2, 0, 2}}
	expect := []int{0, 2}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "ABB"
	queries := [][]int{{2, 0, 2}, {1, 2}, {2, 0, 2}}
	expect := []int{1, 0}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "BABA"
	queries := [][]int{{2, 0, 3}, {1, 1}, {2, 1, 3}}
	expect := []int{0, 1}
	runSample(t, s, queries, expect)
}
