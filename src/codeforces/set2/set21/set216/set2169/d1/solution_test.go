package main

import "testing"

func runSample(t *testing.T, x int, y int, k int, expect int) {
	res := solve(x, y, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 2
	y := 3
	k := 5
	expect := 10
	runSample(t, x, y, k, expect)
}

func TestSample2(t *testing.T) {
	x := 2
	y := 5
	k := 1
	expect := 1
	runSample(t, x, y, k, expect)
}

func TestSample3(t *testing.T) {
	x := 20
	y := 2
	k := 1000000000000
	expect := -1
	runSample(t, x, y, k, expect)
}

func TestSample4(t *testing.T) {
	x := 175
	y := 10
	k := 28
	expect := 2339030304
	runSample(t, x, y, k, expect)
}

func TestSample5(t *testing.T) {
	x := 100000
	y := 998244353
	k := 1999999999
	expect := 2000199999
	runSample(t, x, y, k, expect)
}

func TestSample6(t *testing.T) {
	x := 1
	y := 1
	k := 1
	expect := -1
	runSample(t, x, y, k, expect)
}