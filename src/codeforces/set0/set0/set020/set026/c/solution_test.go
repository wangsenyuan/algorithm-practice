package main

import "testing"

func runSample(t *testing.T, n int, m int, a int, b int, c int, expect bool) {
	res := solve(n, m, a, b, c)
	if (len(res) == n) != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, a, b, c := 2, 6, 2, 2, 1
	expect := true
	runSample(t, n, m, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	n, m, a, b, c := 1, 1, 100, 100, 100

	expect := false
	runSample(t, n, m, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	n, m, a, b, c := 4, 4, 10, 10, 10

	expect := true
	runSample(t, n, m, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	n, m, a, b, c := 3, 2, 1, 0, 1

	expect := true
	runSample(t, n, m, a, b, c, expect)
}

func TestSample5(t *testing.T) {
	n, m, a, b, c := 29, 10, 89, 28, 14

	expect := true
	runSample(t, n, m, a, b, c, expect)
}
