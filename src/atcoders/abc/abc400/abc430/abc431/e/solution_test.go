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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
ABCB
CACC
BCBA`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
CB
AA`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 10
BCBCBCBCBC`
	expect := 10
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 10
CCAABACAAA
CCCBACACCA
BACAABCBBA
ACCCAACCCA
CCAAAACCBA
AACBBACCAA
BCCCACBBAB
CBBCAACCCC
CBBCCBCBCA
BBACABBACC`
	expect := 5
	runSample(t, s, expect)
}