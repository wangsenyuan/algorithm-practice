package main

import "testing"

func runSample(t *testing.T, x int, expect int) {
	res := solve(x)
	if res == expect {
		return
	}
	if expect < 0 || res < 0 {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	tmp := x ^ res
	if res < 1 || res >= x {
		t.Fatalf("Sample result %d exceeds the range", res)
	}
	if tmp == 0 || res+tmp <= x || res+x <= tmp {
		t.Fatalf("Sample result %d, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, -1)
}

func TestSample5(t *testing.T) {
	runSample(t, 69, 66)
}

func TestSample6(t *testing.T) {
	runSample(t, 420, 320)
}

func TestSample7(t *testing.T) {
	// 1011
	runSample(t, 11, 6)
}
