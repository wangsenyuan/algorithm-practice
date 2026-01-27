package main

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, x int, expect bool) {
	res := solve(n, x)
	if len(res) == n != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	var sum int
	for _, v := range res {
		sum ^= v
	}
	if sum != x {
		t.Fatalf("Sample result %v, not correct", res)
	}

	slices.Sort(res)
	res = slices.Compact(res)
	if len(res) != n {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 5, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 6, true)
}

func TestSample3(t *testing.T) {
	// Generate random test cases with n between 100 and 1000
	for i := 0; i < 10; i++ {
		n := rand.Intn(901) + 100 // n in [100, 1000]
		x := rand.Intn(100001)    // x in [0, 100000]
		t.Run(fmt.Sprintf("n=%d_x=%d", n, x), func(t *testing.T) {
			runSample(t, n, x, true)
		})
	}
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 0, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 3, 3, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 3, 0, true)
}

func TestSample7(t *testing.T) {
	runSample(t, 32, 31, true)
}