package main

import "testing"

func runSample(t *testing.T, n int, m int, pies [][]int, expect bool) {
	res := solve(n, m, pies)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 8, 8
	pies := [][]int{
		{4, 4},
		{5, 5},
	}
	expect := true
	runSample(t, n, m, pies, expect)
}
