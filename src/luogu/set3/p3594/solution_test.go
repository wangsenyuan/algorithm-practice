package main

import "testing"

func runSample(t *testing.T, n int, p int, d int, a []int, expect int) {
	res := solve(n, p, d, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p, d := 9, 7, 2
	a := []int{3, 4, 1, 9, 4, 1, 7, 1, 3}
	expect := 5
	runSample(t, n, p, d, a, expect)
}
