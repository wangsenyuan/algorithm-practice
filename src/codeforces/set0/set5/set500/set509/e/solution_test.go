package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	res := solve(s)
	if math.Abs(res-expect)/math.Max(1.0, expect) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "IEAIAIO"
	expect := bruteForce(s)	
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "BYOB"
	expect := bruteForce(s)	
	runSample(t, s, expect)
}