package p3772

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, good []int, expect []int) {
	res := maxSubgraphScore(n, edges, good)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	good := []int{1, 0, 1}
	expect := []int{1, 1, 1}
	runSample(t, n, edges, good, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}, {3, 4}}
	good := []int{0, 1, 0, 1, 1}
	expect := []int{2, 3, 2, 3, 3}
	runSample(t, n, edges, good, expect)
}
