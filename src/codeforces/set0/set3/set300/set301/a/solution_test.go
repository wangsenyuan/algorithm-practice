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
	runSample(t, "2\n50 50 50\n", 150)
}

func TestSample2(t *testing.T) {
	runSample(t, "2\n-1 -100 -1\n", 100)
}

func TestSample3(t *testing.T) {
	runSample(t, "3\n-959 -542 -669 -513 160\n", 2843)
}

func TestOddNCanFlipSingleElement(t *testing.T) {
	runSample(t, "3\n-10 1 1 1 1\n", 14)
}
