package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
4 2
1 4
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `0
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
2 1
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
1 4
`, 5)
}
