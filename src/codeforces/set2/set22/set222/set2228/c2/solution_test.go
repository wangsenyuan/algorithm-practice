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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `0 1
0
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `11 2
1 2
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `222 3
3 4 5
`, 111)
}

func TestSample4(t *testing.T) {
	runSample(t, `3333 4
6 7 8 9
`, 2334)
}

func TestSample5(t *testing.T) {
	runSample(t, `0 1
1
`, 1)
}

func TestDigitsAreProcessedMostSignificantFirst(t *testing.T) {
	runSample(t, `29 2
1 3
`, 2)
}

func TestZeroCannotLeadALongerNumber(t *testing.T) {
	runSample(t, `5 2
0 6
`, 1)
}

func TestSuccessorMayBacktrackIntoMatchedPrefix(t *testing.T) {
	runSample(t, `119 3
0 1 2
`, 1)
}

func TestPredecessorMayBacktrackIntoMatchedPrefix(t *testing.T) {
	runSample(t, `780 3
7 8 9
`, 1)
}

func TestLongerCandidateMustNotOverflowInt64(t *testing.T) {
	runSample(t, `100000000000000000 1
9
`, 1)
}

func TestLongerCandidateMayStartWithNonZeroThenUseZeros(t *testing.T) {
	runSample(t, `999 2
0 1
`, 1)
}
