package p3905

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, sources [][]int, expect [][]int) {
	res := colorGrid(n, m, sources)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	sources := [][]int{{0, 0, 1}, {2, 2, 2}}
	expect := [][]int{{1, 1, 2}, {1, 2, 2}, {2, 2, 2}}

	runSample(t, n, m, sources, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	sources := [][]int{{0, 1, 3}, {1, 1, 5}}

	expect := [][]int{{3, 3, 3}, {5, 5, 5}, {5, 5, 5}}
	runSample(t, n, m, sources, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 2
	sources := [][]int{{1, 1, 5}}
	expect := [][]int{{5, 5}, {5, 5}}
	runSample(t, n, m, sources, expect)
}

func TestSample4(t *testing.T) {
	n, m := 2, 3
	sources := [][]int{{0, 2, 926647}, {1, 0, 485097}}
	expect := [][]int{{485097, 926647, 926647}, {485097, 485097, 926647}}
	runSample(t, n, m, sources, expect)
}
