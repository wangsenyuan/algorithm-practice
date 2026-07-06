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
	runSample(t, `5
10101
60 45 30 40 80
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
000
1 2 3
`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
10101
60 50 50 50 60
`, 4)
}
