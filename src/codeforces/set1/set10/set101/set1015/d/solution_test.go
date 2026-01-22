package main

import (
	"testing"
)

func runSample(t *testing.T, n int, k int, s int, expect bool) {
	res := solve(n, k, s)

	if len(res) == k != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	pos := 1
	var dist int
	for _, x := range res {
		if x == pos {
			t.Fatalf("Sample result %v, not valid, can't stay at the same position", res)
		}
		if x < 1 || x > n {
			t.Fatalf("Sample result %v, not valid, position out of range", res)
		}
		dist += abs(x - pos)
		pos = x
	}
	if dist != s {
		t.Fatalf("Sample result %v, not valid, distance not correct, expect %d, but got %d", res, s, dist)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 2, 15, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 9, 45, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 9, 81, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 9, 82, false)
}