package main

import "testing"

func runSample(t *testing.T, n int, s int, expect int) {
	best, p := solve(n, s)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
	if best < 0 {
		return
	}
	var sum int
	for i, v := range p {
		sum += max(i+1, v)
	}

	if sum != expect {
		t.Fatalf("Sample result %v, not correct, get %d", p, sum)
	}
}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 20, 20)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, 8)
}
