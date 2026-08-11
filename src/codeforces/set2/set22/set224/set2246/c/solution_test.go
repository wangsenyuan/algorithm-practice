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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
-1 1 1 2 3
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
1 3 5 7
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `14
-1 -1 -1 1 2 2 3 3 3 5 5 5 5 5
`, 1536)
}

func TestRegressionOddMinusOneChoicesApplyToEveryAdjacentPair(t *testing.T) {
	runSample(t, `5
-1 -1 1 2 3
`, 6)
}
