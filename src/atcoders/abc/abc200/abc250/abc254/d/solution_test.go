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
	runSample(t, "4\n", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "254\n", 896)
}

func TestSmallDiagonalOnly(t *testing.T) {
	runSample(t, "3\n", 3)
}
