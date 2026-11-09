package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 30
2 1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 13
1 1
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 19
1 2
`, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 50
5 4
`, 1)
}
