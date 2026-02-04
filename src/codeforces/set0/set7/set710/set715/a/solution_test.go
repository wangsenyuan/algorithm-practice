package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	x := 2
	for i := range n {
		add := res[i] * (i + 1)
		x += add
		r := int(math.Sqrt(float64(x)))
		if r*r != x {
			t.Fatalf("Sample result %v, not valid", res)
		}
		if r%(i+2) != 0 {
			t.Fatalf("Sample result %v, not valid", res)
		}
		x = r
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5)
}


func TestSample2(t *testing.T) {
	runSample(t, 1000)
}
