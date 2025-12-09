package main

import "testing"

func runSample(t *testing.T, a int) {
	var time int
	ask := func(x int, y int) string {
		time++
		if time > 60 {
			t.Fatalf("Sample asked too much times %d", time)
		}
		if x%a >= y%a {
			return "x"
		}
		return "y"
	}

	res := solve(ask)

	if res != a {
		t.Fatalf("Sample expect %d, but got %d", a, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 11)
}

func TestSample2(t *testing.T) {
	runSample(t, 1e9)
}

func TestSample3(t *testing.T) {
	runSample(t, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 3)
}
