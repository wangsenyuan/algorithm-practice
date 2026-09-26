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
	runSample(t, `4 2
1 2 3 4
`, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3
5 5 7 3 1
`, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 3
1 2
`, 0)
}
