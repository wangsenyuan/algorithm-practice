package main

import "testing"

func runSample(t *testing.T, m int, a int, b int, expect int) {
	res := solve(m, a, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 4
	a := 1
	b := 2
	expect := 7
	runSample(t, m, a, b, expect)
}

func TestSample2(t *testing.T) {
	m := 1000
	a := 784
	b := 385
	expect := 995373
	runSample(t, m, a, b, expect)
}
