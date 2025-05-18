package p3553

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, edges [][]int, queries [][]int, expect []int) {
	res := minimumWeight(edges, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{0, 1, 2}, {1, 2, 3}, {1, 3, 5}, {1, 4, 4}, {2, 5, 6},
	}
	queries := [][]int{
		{2, 3, 4}, {0, 2, 5},
	}
	expect := []int{12, 11}
	runSample(t, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{
		{1, 0, 8}, {0, 2, 7},
	}
	queries := [][]int{
		{0, 1, 2},
	}
	expect := []int{15}
	runSample(t, edges, queries, expect)
}
