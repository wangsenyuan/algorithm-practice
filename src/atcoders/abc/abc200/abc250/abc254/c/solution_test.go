package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 2
3 4 1 3 4
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3
3 4 1 3 4
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 5
1 2 3 4 5 5 10
`, true)
}
