package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, queries [][]int, expect []int) {
	res := solve(n, m, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	queries := [][]int{
		{1, 1, 2, 1},
		{3, 2, 2},
		{2, 3, 2},
		{3, 3, 3},
		{3, 3, 1},
		{1, 2, 3, 3},
		{3, 3, 2},
		{3, 2, 3},
		{3, 1, 2},
	}
	expect := []int{1, 2, 2, 5, 3, 4}
	runSample(t, n, m, queries, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 1
	queries := [][]int{
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{1, 1, 1, 1000000000},
		{3, 1, 1},
	}
	expect := []int{9000000000}
	runSample(t, n, m, queries, expect)
}

func TestSample3(t *testing.T) {
	n, m := 10, 10
	queries := [][]int{
		{1, 1, 8, 5},
		{2, 2, 6},
		{3, 2, 1},
		{3, 4, 7},
		{1, 5, 9, 7},
		{3, 3, 2},
		{3, 2, 8},
		{2, 8, 10},
		{3, 8, 8},
		{3, 1, 10},
	}
	expect := []int{6, 5, 5, 13, 10, 0}
	runSample(t, n, m, queries, expect)
}
