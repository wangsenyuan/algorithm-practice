package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, res := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}

	// When no swap needed, solution returns [1,1]; swapping index 0 with 0 is a no-op.
	// Verify the (possibly swapped) string never has level < 0.
	buf := []byte(s)
	if res != nil {
		res[0]--
		res[1]--
		buf[res[0]], buf[res[1]] = buf[res[1]], buf[res[0]]
	}
	var level int
	for i := range len(buf) {
		if buf[i] == '+' {
			level++
		} else {
			level--
		}
		if level < 0 {
			t.Fatalf("Sample result %v: level < 0 at %d", res, i)
		}
	}
	// Solution allows more + than - (level > 0 at end is OK)
}

func TestSample1(t *testing.T) {
	runSample(t, "-+", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "+-", true)
}

func TestSample3(t *testing.T) {
	runSample(t, "++--", true)
}

func TestSample4(t *testing.T) {
	runSample(t, "+-+-", true)
}

func TestOddLength(t *testing.T) {
	// Odd length is allowed; only level < 0 makes it impossible
	runSample(t, "++-", true)  // level 1,2,1 — never < 0
	runSample(t, "-", false)   // level -1
	runSample(t, "++", true)   // level 2 — more + than - is OK
}

func TestLevelNotZero(t *testing.T) {
	// More + than - (level > 0 at end) is OK; only level < 0 is impossible
	runSample(t, "+++", true)   // level 3
	runSample(t, "---", false) // level -3
	runSample(t, "++--++", true) // level 2
}

func TestNeedSwap(t *testing.T) {
	runSample(t, "+--+", true)
	runSample(t, "-++-", true)
}

func TestSwapDoesNotFix(t *testing.T) {
	// First '-' and last '+' swap still leaves level < 0 somewhere
	runSample(t, "---+++", false)
}
