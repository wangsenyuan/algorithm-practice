package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := minimumPairRemoval(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 2, 3, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 2}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 3}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-2, 1, 2, -1, -1, -2, -2, -1, -1, 1, 1}
	expect := 10
	runSample(t, a, expect)
}
