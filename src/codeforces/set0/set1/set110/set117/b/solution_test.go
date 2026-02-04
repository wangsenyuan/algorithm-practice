package main

import "testing"

func runSample(t *testing.T, a int, b int, mod int, expect int, expectS1 string) {
	player, s1 := solve(a, b, mod)
	if player != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, player)
	}
	if s1 != expectS1 {
		t.Fatalf("Sample expect %s, but got %s", expectS1, s1)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 7, 2, "")
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 0, 9, 1, "000000001")
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 7, 8, 2, "")
}
