package p3585

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, queries [][]int, expect []int) {
	res := findMedian(n, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	edges := [][]int{{0, 1, 7}}
	queries := [][]int{{1, 0}, {0, 1}}
	expect := []int{0, 1}
	runSample(t, n, edges, queries, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 2}, {2, 0, 4}}
	queries := [][]int{{0, 1}, {2, 0}, {1, 2}}
	expect := []int{1, 0, 2}
	runSample(t, n, edges, queries, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1, 2}, {0, 2, 5}, {1, 3, 1}, {2, 4, 3}}
	queries := [][]int{{3, 4}, {1, 2}}
	expect := []int{2, 2}
	runSample(t, n, edges, queries, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1, 1}, {1, 2, 11}, {1, 3, 1}, {0, 4, 8}}
	queries := [][]int{{0, 3}, {0, 1}}
	expect := []int{1, 1}
	runSample(t, n, edges, queries, expect)
}
