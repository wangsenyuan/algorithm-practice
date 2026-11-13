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
	runSample(t, `5
3 6 9 12 15
`, 36)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
3 7 5 2
`, 1000000006)
}
