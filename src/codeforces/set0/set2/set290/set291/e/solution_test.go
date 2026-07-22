package main

import (
	"bufio"
	"strings"
	"testing"
	"time"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7
1 ab
5 bacaba
1 abacaba
2 aca
5 ba
2 ba
aba
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `7
1 ab
5 bacaba
1 abacaba
2 aca
5 ba
2 ba
bacaba
`, 4)
}

func TestAdversarialBranching(t *testing.T) {
	const prefix = 20000
	const branches = 20000

	parents := make([]int, branches+1)
	edges := make([]string, branches+1)
	parents[0] = 1
	edges[0] = strings.Repeat("a", prefix)
	for i := 1; i <= branches; i++ {
		parents[i] = 2
		edges[i] = "c"
	}
	target := strings.Repeat("a", prefix) + "b"

	start := time.Now()
	if got := solve(branches+2, parents, edges, target); got != 0 {
		t.Fatalf("expect 0, but got %d", got)
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Fatalf("branch transitions took %v; expected near-linear processing", elapsed)
	}
}
