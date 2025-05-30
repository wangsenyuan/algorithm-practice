package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `abab
ab
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abacaba
aba
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abc
ba
`
	expect := false
	runSample(t, s, expect)
}
