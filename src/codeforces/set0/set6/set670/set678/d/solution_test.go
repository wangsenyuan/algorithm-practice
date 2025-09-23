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
	runSample(t, "3 4 1 1", 7)
}

func TestSample2(t *testing.T) {
	runSample(t, "3 4 2 1", 25)
}

func TestSample3(t *testing.T) {
	runSample(t, "3 4 3 1", 79)
}

func TestSample4(t *testing.T) {
	runSample(t, "1 1 1 1", 2)
}

func TestSample5(t *testing.T) {
	runSample(t, "1 192783664 1000000000000000000 596438713", 42838179)
}
