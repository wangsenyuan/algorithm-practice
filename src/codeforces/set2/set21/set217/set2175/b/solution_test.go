package main

import (
	"testing"
)

func runSample(t *testing.T, n int, l int, r int) {
	res := solve(n, l, r)
	l--
	r--
	for i := range n {
		var sum int
		for j := i; j < n; j++ {
			sum ^= res[j]
			if i == l && j == r && sum != 0 {
				t.Fatalf("Sample result %v, is not valid", res)
			} else if (i != l || j != r) && sum == 0 {
				t.Fatalf("Sample result %v, is not valid", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 1, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 2, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 3, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 100000, 3, 4)
}

func TestSample6(t *testing.T) {
	runSample(t, 4, 1, 4)
}


