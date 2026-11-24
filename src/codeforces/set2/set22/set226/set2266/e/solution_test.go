package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1 1\n1\n", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "6 2\n6 6 4 3 2 1\n", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "8 1\n8 6 4 3 2 1 8 6\n", 25)
}

func TestSample4(t *testing.T) {
	runSample(t, "12 3\n12 10 9 8 7 6 5 4 3 2 1 12\n", 15)
}

func TestSample5(t *testing.T) {
	runSample(t, "10 9\n10 9 8 7 6 5 4 3 2 1\n", 1)
}

func TestSample6(t *testing.T) {
	runSample(t, "5 5\n5 4 3 2 1\n", 0)
}

func TestSample7(t *testing.T) {
	runSample(t, "12 4\n12 1 1 1 1 1 1 1 1 1 1 1\n", 1)
}

func TestSample8(t *testing.T) {
	runSample(t, "20 6\n20 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1\n", 1)
}

func TestSample9(t *testing.T) {
	runSample(t, "24 8\n24 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1\n", 1)
}
