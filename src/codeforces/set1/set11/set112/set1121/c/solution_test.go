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
	s := `2 2
49 100
`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `4 2
32 100 33 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `14 5
48 19 6 9 50 20 3 42 38 43 36 21 44 6
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 2
50 100
`
	expect := 0
	runSample(t, s, expect)
}
