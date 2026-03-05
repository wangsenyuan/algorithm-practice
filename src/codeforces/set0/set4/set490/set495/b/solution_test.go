package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Errorf("solve(%d, %d) expect %d, but got %d", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 21, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 9435152, 272, 282)
}

func TestSample3(t *testing.T) {
	// a == b => infinity (encoded as -1)
	runSample(t, 10, 10, -1)
}

func TestABelowB(t *testing.T) {
	runSample(t, 5, 10, 0)
	runSample(t, 0, 1, 0)
}

func TestInfinity(t *testing.T) {
	runSample(t, 0, 0, -1)
	runSample(t, 7, 7, -1)
}

func TestNoSolution(t *testing.T) {
	// a > b but (a-b) has no divisor > b, e.g. a=8, b=5 => diff=3, divisors 1,3; none > 5
	runSample(t, 8, 5, 0)
}

func TestSmall(t *testing.T) {
	// a=6, b=0 => diff=6, divisors > 0: 1,2,3,6 => 4
	runSample(t, 6, 0, 4)
	// a=2, b=1 => diff=1, divisors > 1: none => 0
	runSample(t, 2, 1, 0)
}