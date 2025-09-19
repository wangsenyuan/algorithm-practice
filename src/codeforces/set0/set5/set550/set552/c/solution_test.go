package main

import "testing"

func runSample(t *testing.T, w int, m int, expect bool) {
	res := solve(w, m)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 7, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 99, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 100, 50, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 7, false)
}

func TestSample5(t *testing.T) {
	runSample(t, 10, 999, true)
}
