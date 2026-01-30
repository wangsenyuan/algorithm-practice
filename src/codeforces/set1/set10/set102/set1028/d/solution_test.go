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
	s := `6
ADD 1
ACCEPT 1
ADD 2
ACCEPT 2
ADD 3
ACCEPT 3
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
ADD 1
ADD 2
ADD 3
ACCEPT 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
ADD 1
ADD 2
ADD 3
ADD 4
ADD 5
ACCEPT 3
ACCEPT 5
`
	expect := 0
	runSample(t, s, expect)
}
