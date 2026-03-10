package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, n)
		for j := range n {
			a[i][j] = j
		}
	}

	for _, cur := range res {
		i, l, r := cur[0], cur[1], cur[2]
		i--
		l--
		r--
		slices.Reverse(a[i][l : r+1])
	}
	marked := make([]bool, n)
	for j := range n {
		clear(marked)
		for i := range n {
			if marked[a[i][j]] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			marked[a[i][j]] = true
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 100)
}

func TestSample4(t *testing.T) {
	runSample(t, 201)
}
