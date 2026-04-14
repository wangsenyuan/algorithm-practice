package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// Line 1-2-3-4: only valid pair is path(1,2) & path(3,4) in both orderings
	runSample(t, "4\n1 2\n2 3\n3 4\n", 2)
}

func TestSingleVertex(t *testing.T) {
	runSample(t, "1\n", 0)
}

func TestTwoVertices(t *testing.T) {
	runSample(t, "2\n1 2\n", 0)
}

func TestThreeVertices(t *testing.T) {
	// Line 1-2-3: paths (1,2),(2,3) share 2; (1,3) covers all; no valid pair
	runSample(t, "3\n1 2\n2 3\n", 0)
}

func TestStar(t *testing.T) {
	// Star: center=1, leaves 2,3,4,5 (n=5)
	// Paths of length 1: (1,2),(1,3),(1,4),(1,5) — all share center
	// Paths of length 2: (2,3),(2,4),(2,5),(3,4),(3,5),(4,5) — all go through center
	// Any two paths share center 1 → answer = 0
	runSample(t, "5\n1 2\n1 3\n1 4\n1 5\n", 0)
}
