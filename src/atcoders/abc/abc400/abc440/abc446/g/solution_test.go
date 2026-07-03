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
	runSample(t, `8
2 1 2 1 1 2 7 2
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
2 3 4 5 4
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `20
2 2 3 1 1 4 1 4 1 4 2 4 1 2 1 4 4 1 1 4
`, 15)
}
