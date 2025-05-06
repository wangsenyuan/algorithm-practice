package p3537

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect [][]int) {
	res := specialGrid(n)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 0
	expect := [][]int{
		{0},
	}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	expect := [][]int{
		{3, 0},
		{2, 1},
	}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	expect := [][]int{
		{15, 12, 3, 0},
		{14, 13, 2, 1},
		{11, 8, 7, 4},
		{10, 9, 6, 5},
	}
	runSample(t, n, expect)
}
