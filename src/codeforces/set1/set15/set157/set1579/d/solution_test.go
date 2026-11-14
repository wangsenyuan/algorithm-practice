package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	got := drive(reader)
	if len(got) != expect {
		t.Fatalf("Sample expect %d times meeting, but got %d", expect, len(got))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3\n1 2 3\n", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "4\n1 0 1 0\n", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "2\n1 1\n", 1)
}
