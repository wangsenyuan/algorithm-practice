package main

import "testing"

func runSample(t *testing.T, vp, vd, tt, f, c int, expect int) {
	ans := solve(vp, vd, tt, f, c)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 1, 1, 10, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 1, 1, 8, 1)
}

func TestNoChase(t *testing.T) {
	runSample(t, 5, 5, 3, 1, 100, 0)
}

func TestReachBeforeNotice(t *testing.T) {
	runSample(t, 10, 100, 1, 5, 5, 0)
}

func TestSingleBijouExactCastle(t *testing.T) {
	runSample(t, 1, 2, 1, 1, 4, 1)
}
