package p3515

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := treeQueries(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{
		{1, 2, 7},
	}
	queries := [][]int{
		{2, 2}, {1, 1, 2, 4}, {2, 2},
	}
	expect := []int{7, 4}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{
		{1, 2, 2}, {1, 3, 4},
	}
	queries := [][]int{
		{2, 1}, {2, 3}, {1, 1, 3, 7}, {2, 2}, {2, 3},
	}
	expect := []int{0, 4, 2, 7}
	runSample(t, n, edges, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 2, 2}, {2, 3, 1}, {3, 4, 5},
	}
	queries := [][]int{
		{2, 4}, {2, 3}, {1, 2, 3, 3}, {2, 2}, {2, 3},
	}
	expect := []int{8, 3, 2, 5}
	runSample(t, n, edges, queries, expect)
}
