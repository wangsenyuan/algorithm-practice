package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s + "\n")))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "ouuokarinn", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "banana", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "nnn", 3)
}

func TestSample4(t *testing.T) {
	runSample(t, "amanda", 0)
}
