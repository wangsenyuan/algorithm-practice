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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
0 2 1`

	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4
0 2`

	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
0 0 1 1`

	expect := 8
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 7
6 6 2 4 3 0 1 8`

	expect := 25
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 2
0 0`

	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 4
0 0`

	expect := 0
	runSample(t, s, expect)
}
