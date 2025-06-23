package p3590

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, par []int, vals []int, queries [][]int, expect []int) {
	res := kthSmallest(par, vals, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	par := []int{-1, 0, 0}
	vals := []int{1, 1, 1}
	queries := [][]int{{0, 1}, {0, 2}, {0, 3}}
	expect := []int{0, 1, -1}
	runSample(t, par, vals, queries, expect)
}

func TestSample2(t *testing.T) {
	par := []int{-1, 0, 1}
	vals := []int{5, 2, 7}
	queries := [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 1}}
	expect := []int{0, 7, -1, 0}
	runSample(t, par, vals, queries, expect)
}
