package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func prefixMexXor(p []int) int {
	n := len(p)
	seen := make([]bool, n+2)
	mex := 0
	x := 0
	for _, v := range p {
		if v >= 0 && v < len(seen) {
			seen[v] = true
		}
		for mex < len(seen) && seen[mex] {
			mex++
		}
		x ^= mex
	}
	return x
}

func isPermutation(p []int) bool {
	n := len(p)
	seen := make([]bool, n)
	for _, v := range p {
		if v < 0 || v >= n || seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func runSample(t *testing.T, s string, expectOK bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var n, k int
	fmt.Fscan(reader, &n, &k)

	ok, p := drive(bufio.NewReader(strings.NewReader(s)))
	if ok != expectOK {
		t.Fatalf("Sample expect ok=%v, but got ok=%v p=%v", expectOK, ok, p)
	}
	if !ok {
		return
	}
	if len(p) != n || !isPermutation(p) {
		t.Fatalf("Sample expect a permutation of 0..%d-1, but got %v", n, p)
	}
	if got := prefixMexXor(p); got != k {
		t.Fatalf("Sample permutation %v has prefix-mex xor %d, want %d", p, got, k)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 0
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 0
`, true)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 8
`, false)
}

func TestSample5(t *testing.T) {
	runSample(t, `5 1
`, true)
}

func TestSample6(t *testing.T) {
	runSample(t, `9 12
`, true)
}
