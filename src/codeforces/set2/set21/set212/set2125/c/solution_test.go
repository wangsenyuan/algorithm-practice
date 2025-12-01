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
	s := "2 100"
	runSample(t, s, 21)
}

func TestSample2(t *testing.T) {
	s := "2 1000"
	runSample(t, s, 227)
}

func TestSample3(t *testing.T) {
	s := "13 37"
	runSample(t, s, 7)
}

func TestSample4(t *testing.T) {
	s := "2 1000000000000000000"
	runSample(t, s, 228571428571428570)
}
