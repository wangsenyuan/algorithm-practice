package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
1011`, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
10`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `31
1000000010111001111000111001110`, 129377)
}

func TestSample4(t *testing.T) {
	runSample(t, `62
00010011000110010011110110011001110110010011110110111100100010`, 996654969)
}

// func TestSample5(t *testing.T) {
// 	runSample(t, `75
// 011001100010010010100010011010001000110010011010100111110110100000010111111`, 928344407)
// }
