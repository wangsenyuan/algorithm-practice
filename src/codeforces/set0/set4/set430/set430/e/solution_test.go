package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 1 1 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 1 5 2 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `19
9 7 1 8 1 1 1 13 1 1 3 3 19 1 1 1 1 1 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
1 2
`
	expect := false
	runSample(t, s, expect)
}