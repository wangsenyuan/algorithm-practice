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
	s := `5
-1 -1 4 3 -1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
-1 -1 4 -1 -1
`
	// f(4) = 3 * 2 = 6
	// f(5) = 4 * 6 = 24

	expect := 11
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
-1 -1 4 -1 7 1 6
`

	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
-1 -1 -1 -1 -1 -1
`

	expect := 265
	runSample(t, s, expect)
}
