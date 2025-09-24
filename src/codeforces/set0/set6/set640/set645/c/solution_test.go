package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 2
0100100`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `5 1
01010`
	runSample(t, s, 2)
}

func TestSample3(t *testing.T) {
	s := `3 2
000`
	runSample(t, s, 1)
}

func TestSample4(t *testing.T) {
	s := `112 12
0110101000000010101110010111100101011010011110100111111100011101011111000111101101110100111011110001100110110010`
	runSample(t, s, 10)
}
