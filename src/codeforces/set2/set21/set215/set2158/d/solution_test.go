package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, y string, expect bool) {
	ok, res := solve(s, y)

	if ok != expect {
		t.Fatalf("Sample result %v, got %t, not expected %t", res, ok, expect)
	}
	if !ok {
		return
	}
	n := len(s)

	if len(res) > 2*n {
		t.Fatalf("Sample result %v, took too much steps", res)
	}
	a := convert(s)
	b := convert(y)

	flip := func(l int, r int) {
		for l < r {
			if a[l] != a[r] {
				t.Fatalf("Sample result %v, got %v, not expected %v", res, a, b)
			}
			a[l] ^= 1
			a[r] ^= 1
			l++
			r--
		}

		if l == r {
			a[l] ^= 1
		}
	}

	for _, cur := range res {
		flip(cur[0]-1, cur[1]-1)
	}

	if !slices.Equal(a, b) {
		t.Fatalf("Sample result %v, got %v, not expected %v", res, a, b)
	}
}

func TestSample1(t *testing.T) {
	s := "01011"
	y := "10000"
	runSample(t, s, y, true)
}

func TestSample2(t *testing.T) {
	s := "1010101"
	y := "0101010"
	runSample(t, s, y, true)
}

func TestSample3(t *testing.T) {
	s := "0101"
	y := "1000"
	runSample(t, s, y, true)
}

func TestSample4(t *testing.T) {
	s := "0101"
	y := "1010"
	runSample(t, s, y, true)
}
