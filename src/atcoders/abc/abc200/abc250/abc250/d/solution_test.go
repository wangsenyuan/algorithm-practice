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
	runSample(t, `250
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `123456789012345
`, 226863)
}
