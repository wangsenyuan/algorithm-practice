package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
4 0 2
3 5 7
8 1 6
`, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 1 1 1
1 1 0 1
1 1 1 1
1 1 1 1
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
1 1 1 1
1 1 0 1
1 1 2 1
1 1 1 1
`, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
1 2 2
2 2 1
0 1 2
`, -1)
}
