package main

import (
	"math/rand"
	"slices"
	"testing"
)

func runSample(t *testing.T, n int) {
	res := solve(n)
	if len(res) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	var bad int
	for i := 0; i+2 < n; i++ {
		if gcd(res[i], res[i+1]) == 1 && gcd(res[i+1], res[i+2]) == 1 && gcd(res[i], res[i+2]) == 1 {
			bad++
		}
	}

	if bad > 6 {
		t.Fatalf("Sample result %v, has too many bad positions %d", res, bad)
	}
	slices.Sort(res)
	res = slices.Compact(res)

	if len(res) != n || res[0] != 1 || res[n-1] != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 100)
}

func TestSample2(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for range 100 {
		n := rng.Intn(9999-1001+1) + 1001
		runSample(t, n)
	}
}

func TestSample3(t *testing.T) {
	runSample(t, 6)
}
