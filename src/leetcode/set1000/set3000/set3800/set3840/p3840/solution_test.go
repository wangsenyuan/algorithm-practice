package p3840

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, s string, queries []string, expect []bool) {
	res := palindromePath(n, edges, s, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	s := "aac"
	queries := []string{"query 0 2", "update 1 b", "query 0 2"}
	expect := []bool{true, false}
	runSample(t, n, edges, s, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}}
	s := "abca"
	queries := []string{"query 1 2", "update 0 b", "query 2 3", "update 3 a", "query 1 3"}
	expect := []bool{false, false, true}
	runSample(t, n, edges, s, queries, expect)
}
