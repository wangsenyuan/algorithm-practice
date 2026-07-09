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
	runSample(t, `6 4
1 6 9 2 3 1
1 10 3 7
`, 508)
}

func TestSample2(t *testing.T) {
	runSample(t, `20 20
36625 195265 98908 111868 111868 47382 147644 472464 472464 416653 111868 195265 327972 327972 262769 75439 381156 451275 36625 195265
327972 111868 416653 177330 340019 262769 47382 262769 47382 340019 47382 262769 327972 327972 359676 381156 327972 36625 451275 381156
`, 58141644)
}
