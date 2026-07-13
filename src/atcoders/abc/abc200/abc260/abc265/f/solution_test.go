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
	runSample(t, `1 5
0
3
`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 10
2 6 5
2 1 2
`, 632)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 100
3 1 4 1 5 9 2 6 5 3
2 7 1 8 2 8 1 8 2 8
`, 145428186)
}
