package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectX int, expectY int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	x, y := drive(reader)
	if x != expectX || y != expectY {
		t.Fatalf("Sample expect (%d, %d), but got (%d, %d)", expectX, expectY, x, y)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3\n", -2, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "7\n", 3, 2)
}

func TestZeroMoves(t *testing.T) {
	runSample(t, "0\n", 0, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "14\n", -2, -4)
}

func TestSample4(t *testing.T) {
	runSample(t, "59\n", 7, -2)
}

func TestSample5(t *testing.T) {
	runSample(t, "60\n", 8, 0)
}

func TestSample6(t *testing.T) {
	runSample(t, "1000000000000000000\n", -418284973, -1154700538)
}
