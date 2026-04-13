package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1\n2\n", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "2\n1 2\n", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "2\n3 5\n", 1)
}

// 3-4-5 primitive triple: all three lawns in one component.
func TestTriangle345(t *testing.T) {
	runSample(t, "3\n3 4 5\n", 1)
}

// 5-12-13: one component.
func TestTriangle51213(t *testing.T) {
	runSample(t, "3\n5 12 13\n", 1)
}

// Isolated values plus one triple component {3,4,5}.
func TestMixedIsolatedAndTriple(t *testing.T) {
	runSample(t, "5\n1 2 3 4 5\n", 3)
}

// Scaled 3-4-5 is not pairwise coprime → 6 and 8 are not adjacent via that triple alone;
// still expect a single component if another primitive triple links them (here 6-8-10 fails beautiful).
// 6 and 8: no beautiful triple with integer third side? 6^2+8^2=10^2 but gcd(6,8,10)≠1 pairwise.
// So {6,8} should be 2 components when alone.
func TestScaledTripleNotBeautiful(t *testing.T) {
	runSample(t, "2\n6 8\n", 2)
}

// Single lawn.
func TestSingle(t *testing.T) {
	runSample(t, "1\n9999999\n", 1)
}
