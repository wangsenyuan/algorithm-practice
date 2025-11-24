package p3756

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, queries [][]int, expect []int) {
	res := sumAndMultiply(s, queries)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10203004"
	queries := [][]int{
		{0, 7}, {1, 3}, {4, 6},
	}
	expect := []int{12340, 4, 9}
	runSample(t, s, queries, expect)
}

func TestSample2(t *testing.T) {
	s := "1000"
	queries := [][]int{
		{0, 3}, {1, 1},
	}
	expect := []int{1, 0}
	runSample(t, s, queries, expect)
}

func TestSample3(t *testing.T) {
	s := "9876543210"
	queries := [][]int{
		{0, 9},
	}
	expect := []int{444444137}
	runSample(t, s, queries, expect)
}
