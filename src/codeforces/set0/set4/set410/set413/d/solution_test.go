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
	s := `7 4
2 2 4 2 2 2 2`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `1 3
0`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `2 3
0 4`
	runSample(t, s, 1)
}

func TestSample4(t *testing.T) {
	s := `5 4
2 0 0 4 4`
	runSample(t, s, 2)
}

func TestSample5(t *testing.T) {
	s := `6 4
4 2 0 4 4 0`
	// 4 2 2 4 4 4
	// 4 2 2 4 4 2
	// 4 2 4 4 4 4
	runSample(t, s, 3)
}

func TestSample6(t *testing.T) {
	s := `5 3
4 4 0 0 0`
	runSample(t, s, 8)
}
