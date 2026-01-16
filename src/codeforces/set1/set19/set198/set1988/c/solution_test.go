package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
	for i, v := range res {
		if v > n || v <= 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		if i > 0 && v|res[i-1] != n {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	if !sort.IntsAreSorted(res) {
		t.Fatalf("Sample result %v, is not sorted", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 14, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 23, 5)
}

func TestSample6(t *testing.T) {
	runSample(t, 23, 5)
}

func TestSample7(t *testing.T) {
	runSample(t, 2, 1)
}