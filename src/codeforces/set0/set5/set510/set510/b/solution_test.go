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
	s := `3 4
AAAA
ABCA
AAAA`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
AAAA
ABCA
AADA`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
YYYR
BYBY
BBBY
BBBY`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7 6
AAAAAB
ABBBAB
ABAAAB
ABABBB
ABAAAB
ABBBAB
AAAAAB`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 13
ABCDEFGHIJKLM
NOPQRSTUVWXYZ
`
	expect := false
	runSample(t, s, expect)
}
