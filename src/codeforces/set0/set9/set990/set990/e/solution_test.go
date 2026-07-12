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
	runSample(t, `6 2 3
1 3
1 2 3
`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3 4
1 2 3
1 10 100 1000
`, 1000)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 1 5
0
3 3 3 3 3
`, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, `7 4 3
2 4 5 6
3 14 15
`, -1)
}

func TestCanMoveToNearestUsableBeforeBlockedEndpoint(t *testing.T) {
	runSample(t, `3 1 2
2
10 1
`, 2)
}
