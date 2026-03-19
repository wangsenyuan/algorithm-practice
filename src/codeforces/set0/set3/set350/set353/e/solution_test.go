package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "001", 1)
}

func TestSample2(t *testing.T) {
	// 0 <- 1 <- 2 -> 3 -> 4 <- 5 -> 6 -> 0
	runSample(t, "110010", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "0000100000110010100010010100111001000111000101101001101010110001001010111010111011000011111110001010", 34)
}

func TestSample4(t *testing.T) {
	runSample(t, "0111001010101110001100000010011000100101110010001100111110101110001110101010111000010010011000000110", 35)
}
