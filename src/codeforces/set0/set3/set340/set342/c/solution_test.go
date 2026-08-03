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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 1
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 2
`, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 1
`, 2)
}

func TestRemainderCases(t *testing.T) {
	runSample(t, `10 3
`, 1)
	runSample(t, `10 5
`, 2)
	runSample(t, `10 9
`, 3)
	runSample(t, `10 13
`, 3)
}
