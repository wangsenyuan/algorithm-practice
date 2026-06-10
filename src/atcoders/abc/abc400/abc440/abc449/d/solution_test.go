package main

import "testing"

func runSample(t *testing.T, L int, R int, D int, U int, expect int) {
	res := solve(L, R, D, U)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L, R, D, U := -4, 3, 1, 3
	expect := 10
	runSample(t, L, R, D, U, expect)
}

func TestSample2(t *testing.T) {
	L, R, D, U := -14, 14, -14, 14
	expect := 449
	runSample(t, L, R, D, U, expect)
}
