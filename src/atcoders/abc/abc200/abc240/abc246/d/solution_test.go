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
	runSample(t, "9\n", 15)
}

func TestSample2(t *testing.T) {
	runSample(t, "0\n", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "999999999989449206\n", 1_000_000_000_000_000_000)
}

func TestCubeWithZeroEndpoint(t *testing.T) {
	runSample(t, "1\n", 1)
	runSample(t, "8\n", 8)
	runSample(t, "27\n", 27)
}
