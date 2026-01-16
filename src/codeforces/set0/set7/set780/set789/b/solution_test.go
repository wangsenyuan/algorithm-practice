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
	s := `3 2 30 4
6 14 25 48
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `123 1 2143435 4
123 11 -5453 141245
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `123 1 2143435 4
54343 -13 6 124
`
	expect := -1
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `123 0 21 4
543453 -123 6 1424
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `123 0 21 4
543453 -123 6 1424
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 0 2 1
2
`
	expect := -1
	runSample(t, s, expect)
}
