package main

import (
	"slices"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect bool) {
	res := solve(n, k)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	var sum int
	for i := 0; i+1 < n; i += 2 {
		if res[i] > 1e9 || res[i+1] > 1e9 {
			t.Fatalf("Sample resut %v, should be less than 1e9", res)
		}
		sum += gcd(res[i], res[i+1])
	}
	if sum != k {
		t.Fatalf("Sample expect %d, but got %d", k, sum)
	}

	sort.Ints(res)
	res = slices.Compact(res)

	if len(res) != n {
		t.Fatalf("Sample resut %v, should be distinct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 3, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 100000, 100000000, true)
}
