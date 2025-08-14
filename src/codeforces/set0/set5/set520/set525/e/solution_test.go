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
	s := `2 2 30
4 3
`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `2 2 7
4 3
`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `3 1 1
1 1 1
`
	runSample(t, s, 6)
}

func TestSample4(t *testing.T) {
	s := `25 25 25
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
`
	runSample(t, s, 33554432)
}

