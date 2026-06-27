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
	runSample(t, "21\n", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "407\n", 17)
}

func TestSample3(t *testing.T) {
	runSample(t, "2025524202552420255242025524\n", 150)
}
