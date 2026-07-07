package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	t.Helper()
	res := solve(n)
	if expect == nil {
		if res != nil {
			t.Fatalf("Sample expect no solution, but got %v", res)
		}
		return
	}
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	checkPermutation(t, n, res)
}

func checkPermutation(t *testing.T, n int, p []int) {
	t.Helper()
	if len(p) != n {
		t.Fatalf("permutation length %d, want %d", len(p), n)
	}
	seen := make([]bool, n+1)
	for _, x := range p {
		if x < 1 || x > n || seen[x] {
			t.Fatalf("invalid permutation %v", p)
		}
		seen[x] = true
	}
	for i := 1; i < n; i++ {
		ok := false
		for j := i; j <= n; j++ {
			if p[i-1] == (p[j-1] ^ i) {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("permutation %v fails at i=%d", p, i)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, []int{2, 1, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, 4, nil)
}

func TestSmallSizes(t *testing.T) {
	for n := 3; n <= 64; n++ {
		res := solve(n)
		if n&(n-1) == 0 {
			if res != nil {
				t.Fatalf("n=%d should have no solution, but got %v", n, res)
			}
			continue
		}
		if res == nil {
			t.Fatalf("n=%d should have a solution", n)
		}
		checkPermutation(t, n, res)
	}
}
