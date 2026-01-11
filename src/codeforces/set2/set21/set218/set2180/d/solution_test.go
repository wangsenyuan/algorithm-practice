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
	s := `3
1 2 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 4 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 2 11 12 21 22`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
0 1 2 3 5 8 13`
	expect := 6
	runSample(t, s, expect)
}
