package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, c []int, d [][]int, s []int, expect []int) {
	res := solve(c, d, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{1, 2, 5, 10}
	d := [][]int{
		{3, 2, 3, 1},
		{1000, 2, 2, 2},
	}
	s := []int{10, 900}
	expect := []int{4, 27}
	runSample(t, c, d, s, expect)
}
