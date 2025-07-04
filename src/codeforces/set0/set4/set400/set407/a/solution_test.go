package main

import "testing"

func runSample(t *testing.T, a int, b int, expect bool) {
	res := solve(a, b)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	var origin point

	d1 := origin.dist(point{res[0][0], res[0][1]})
	d2 := origin.dist(point{res[1][0], res[1][1]})

	if d1 != a*a || d2 != b*b {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 10, true)
}