package main

import "testing"

func runSample(t *testing.T, s string, expect byte) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %c, but got %c", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "DDRRR", 'D')
}
func TestSample2(t *testing.T) {
	runSample(t, "DDRRRR", 'R')
}
