package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x int, y int, n int, expect []int) {
	res := solve(x, y, n)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 7, 6, []int{2, 5})
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, 4, []int{7, 2})
}
