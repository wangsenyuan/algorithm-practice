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
	runSample(t, "8\n", 80)
}

func TestSample2(t *testing.T) {
	runSample(t, "461\n", 1385)
}

func TestSample3(t *testing.T) {
	runSample(t, "100\n", 1702)
}
