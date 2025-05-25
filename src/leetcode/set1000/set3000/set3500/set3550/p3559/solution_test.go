package p3559

import "testing"

func runSample(t *testing.T, edges [][]int, queries [][]int, expect []int) {
	res := assignEdgeWeights(edges, queries)

	for i, expect := range expect {
		if res[i] != expect {
			t.Errorf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{
		{1, 2},
	}
	queries := [][]int{
		{1, 1},
		{1, 2},
	}
	expect := []int{0, 1}
	runSample(t, edges, queries, expect)
}
