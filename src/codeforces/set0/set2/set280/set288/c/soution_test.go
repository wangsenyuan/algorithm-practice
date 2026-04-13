package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	sum, p := solve(n)
	if sum != expect || len(p) != n+1 {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
	var tot int
	for i := range n + 1 {
		tot += i ^ p[i]
	}

	if tot != expect {
		t.Fatalf("Sample result %v, not correct, getting %d", p, tot)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 20)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 30)
}

