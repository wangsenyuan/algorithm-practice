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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 16
10 9
10 5 1
12 2 1
20 18 3
10 2 3
2 1 5
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 10
3 1
3 1 1
3 1 2
3 1 3
3 1 4
`, 5)
}

func TestLightChildUsesItsOwnDiscount(t *testing.T) {
	runSample(t, `3 10
10 1
10 1 1
10 9 1
`, 2)
}

func TestLightChildKeepsMaximumCountState(t *testing.T) {
	runSample(t, `3 3
10 9
10 9 1
10 9 1
`, 3)
}
