package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3, 2, 4, 6, 8, 2, 2, 3, 7}
	expect := 62
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 1}
	// a[1], a[1], a[1]
	// a[2], a[2], a[2]
	// a[1], a[2], a[2]
	// a[2], a[1], a[1]
	expect := 8
	runSample(t, a, expect)
}
