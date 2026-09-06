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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
3 1
4 2
`, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
5 3 1
2 4 6
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
7 5 9 1 3
2 4 6 10 8
`, 3)
}
