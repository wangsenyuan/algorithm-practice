package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "(((#)((#)", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "()((#((#(#()", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "(#)", false)
}
