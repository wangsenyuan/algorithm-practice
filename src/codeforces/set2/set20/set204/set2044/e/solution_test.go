package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// (2, 2), (3, 3), (4, 4), (5, 5), (6, 6)
	// (2, 4), (3, 6), (4, 8),  (5, 10), (6, 12)
	// (2, 8), (3, 12)
	runSample(t, "2 2 6 2 12", 12)
}

func TestSample2(t *testing.T) {
	runSample(t, "2 1 1000000000 1 1000000000", 1999999987)
}

func TestSample3(t *testing.T) {
	runSample(t, "3 5 7 15 63", 6)
}

func TestSample4(t *testing.T) {
	runSample(t, "1000000000 1 5 6 1000000000", 1)
}

func TestSample5(t *testing.T) {
	runSample(t, "15 17 78 2596 20914861", 197)
}
