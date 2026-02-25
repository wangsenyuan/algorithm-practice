package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 25, 38, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 12, 20, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 7, 9, 1)
}

func TestSample5(t *testing.T) {
	l, r := 7, 49
	expect := bruteForce(l, r)
	runSample(t, l, r, expect)
}

func TestSample6(t *testing.T) {
	l, r := 7, 1000
	expect := bruteForce(l, r)
	runSample(t, l, r, expect)
}
