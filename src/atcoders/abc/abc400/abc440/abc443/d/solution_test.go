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
	runSample(t, `5
5 2 1 3 4`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
1 1`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
1 3 1`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `9
9 9 8 2 4 4 3 5 3`, 16)
}

func TestSample5(t *testing.T) {
	runSample(t, `20
7 4 6 2 15 5 17 15 1 8 18 1 5 1 12 11 2 7 8 14`, 105)
}
