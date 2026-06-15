package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a, b, expect := 1, 1, 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b, expect := 3, 2, 3
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b, expect := 199, 200, 200
	runSample(t, a, b, expect)
}
