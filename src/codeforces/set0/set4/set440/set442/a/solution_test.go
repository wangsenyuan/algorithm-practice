package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	cards := strings.Split(s, " ")
	res := solve(cards)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "Y4 B1 R3 G5 R5 W3 W5 W2 R1 Y1", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "G3 G3", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "G4 R4 R3 B3", 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "B1 Y1 W1 G1 R1", 4)
}
