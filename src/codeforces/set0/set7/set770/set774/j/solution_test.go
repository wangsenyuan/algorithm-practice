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
	s := `5 2
NYNNY`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
????NN`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 8
NYNNY?YNNNNNN?NNNNNYNY?YYNYNN?NNNY??NNYNYNNNYNNNYNNNNNNNNY?NNNYNYN?NNNY?YY?NNYNN?NNNYNNYNNYN?NNYNYNN`
	expect := true
	runSample(t, s, expect)
}
