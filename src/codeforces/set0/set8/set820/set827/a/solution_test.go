package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
a 4 1 3 5 7
ab 2 1 5
ca 1 4
`
	expect := "abacaba"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
a 1 3
`
	expect := "aaa"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
ab 1 1
aba 1 3
ab 2 3 5
`
	expect := "ababab"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
ba 2 16 18
a 1 12
b 3 4 13 20
bb 2 6 8
ababbbbbaab 1 3
abababbbbb 1 1
`
	expect := "abababbbbbaabaababab"
	runSample(t, s, expect)
}
