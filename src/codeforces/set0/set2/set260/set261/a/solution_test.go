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
	runSample(t, `1
2
4
50 50 100 100
`, 200)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
2 3
5
50 50 50 50 50
`, 150)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
1
7
1 1 1 1 1 1 1
`, 3)
}
