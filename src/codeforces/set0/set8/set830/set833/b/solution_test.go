package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2, 1}
	k := 1
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 3, 1, 4, 4, 4}
	k := 2
	expect := 5
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 7, 8, 7, 7, 8, 1, 7}
	k := 3
	expect := 6
	runSample(t, a, k, expect)
}
