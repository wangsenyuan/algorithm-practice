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
		t.Errorf("expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
4 2 3 1`
	runSample(t, s, 13)
}

func TestSample2(t *testing.T) {
	s := `7
7 6 1 2 3 3 2`
	runSample(t, s, 73)
}

func TestSample3(t *testing.T) {
	s := `5
1 1 1 1 1`
	runSample(t, s, 32)
}

func TestSample4(t *testing.T) {
	s := `11
7 2 1 9 7 3 4 1 3 5 3`
	runSample(t, s, 619)
}
