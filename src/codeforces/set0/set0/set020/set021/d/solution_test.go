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
	runSample(t, `3 3
1 2 1
2 3 1
3 1 1
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 2 3
2 3 4
`, 14)
}

func TestDisconnectedEdge(t *testing.T) {
	runSample(t, `4 2
1 2 5
3 4 7
`, -1)
}

func TestLoopAtReachableVertex(t *testing.T) {
	runSample(t, `2 2
1 2 3
2 2 5
`, 11)
}
