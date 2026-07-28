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
	runSample(t, "1\n", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "2\n", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "3\n", 15)
}

func TestFourthUnsolvableValue(t *testing.T) {
	runSample(t, "4\n", 63)
}
