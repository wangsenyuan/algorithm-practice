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
4 1 2 5 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 1 3 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
1 3 5 7 2 4 6
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
6 2 1 8 5 7 3 4
`
	expect := 5
	runSample(t, s, expect)
}
