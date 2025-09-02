package main

import "testing"

func runSample(t *testing.T, p int, s string, expect int) {
	res := solve(p, s)
	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", p, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := 3
	s := "aeabcaez"
	expect := 6
	runSample(t, p, s, expect)
}

func TestSample2(t *testing.T) {
	p := 4
	s := "rkoa"
	expect := 14
	runSample(t, p, s, expect)
}
