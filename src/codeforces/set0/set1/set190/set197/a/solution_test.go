package main

import "testing"

func runSample(t *testing.T, a int, b int, r int, expect string) {
	res := solve(a, b, r)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 5, 2, "First")
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 7, 4, "Second")
}
