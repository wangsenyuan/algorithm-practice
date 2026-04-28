package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 3
1 1 2 2 3 3
4 5 1
6 7 1
2 3 1`, 2)
}

func TestNoExtraEdges(t *testing.T) {
	runSample(t, `3 0
1 2`, 0)
}

func TestChooseBestOverlappingPath(t *testing.T) {
	runSample(t, `4 3
1 2 3
1 3 5
2 4 7
1 4 10`, 10)
}

func TestCanChooseDisjointPaths(t *testing.T) {
	runSample(t, `7 3
1 1 2 2 3 3
4 5 4
6 7 5
4 7 100`, 100)
}
