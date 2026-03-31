package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 3, 1, 3
	expect := 2
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 3, 3, 2
	expect := 3
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 4, 2, 2
	expect := 3
	runSample(t, n, m, k, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 3, 2, 1
	expect := 2
	runSample(t, n, m, k, expect)
}

func TestSample5(t *testing.T) {
	n, m, k := 4, 3, 3
	expect := 3
	runSample(t, n, m, k, expect)
}

func TestSample6(t *testing.T) {
	n, m, k := 7, 7, 4
	expect := 6
	runSample(t, n, m, k, expect)
}

func TestSample7(t *testing.T) {
	n, m, k := 100000, 1000000000, 100000
	expect := 100000
	runSample(t, n, m, k, expect)
}
