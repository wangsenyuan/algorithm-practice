package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 2, 4, 8, 16, 18
	// 4, 9, 16
	runSample(t, "20\n", 5)
}

func TestSample2(t *testing.T) {
	runSample(t, "400\n", 24)
}

func TestSample3(t *testing.T) {
	runSample(t, "1234567890\n", 42413)
}

func TestSample4(t *testing.T) {
	runSample(t, "4\n", 2)
}

