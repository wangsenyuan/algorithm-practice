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
	runSample(t, `6 4
5237
2753
7523
5723
5327
2537
`, 2700)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
010
909
012
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 5
50808
36603
37198
44911
29994
42543
50156
`, 20522)
}
