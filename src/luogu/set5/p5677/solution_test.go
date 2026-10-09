package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, want int64) {
	t.Helper()
	got := drive(bufio.NewReader(strings.NewReader(input)))
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3 2\n2 1 3\n1 2\n1 3\n", 10)
}

func TestSample2(t *testing.T) {
	runSample(t, "2 1\n5 9\n1 2\n", 2)
}
