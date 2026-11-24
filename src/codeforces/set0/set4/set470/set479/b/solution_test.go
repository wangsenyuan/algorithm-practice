package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	s, ops := drive(bufio.NewReader(strings.NewReader(input)))
	if s != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, s)
	}
	if len(ops) > k {
		t.Fatalf("Sample used %d operations, more than k = %d", len(ops), k)
	}
	for _, op := range ops {
		i, j := op[0]-1, op[1]-1
		if i == j || i < 0 || i >= n || j < 0 || j >= n || a[i] == 0 {
			t.Fatalf("Sample operation %v is not valid", op)
		}
		a[i]--
		a[j]++
	}
	if got := slices.Max(a) - slices.Min(a); got != s {
		t.Fatalf("Sample operations give instability %d, but reported %d", got, s)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
5 8 5
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
2 2 4
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 3
8 3 2 6 3
`, 3)
}
