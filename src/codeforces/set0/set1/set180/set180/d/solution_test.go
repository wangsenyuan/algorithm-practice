package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `aad
aac
`
	expect := "aad"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abad
bob
`
	expect := "daab"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abc
defg
`
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `czaaab
abcdef
`
	expect := "abczaa"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `abacaba
aba
`
	expect := "abaaabc"
	runSample(t, s, expect)
}
