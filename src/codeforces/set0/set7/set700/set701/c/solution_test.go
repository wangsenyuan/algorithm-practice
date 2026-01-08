package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7
bcAAcbc
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `3
AaA
`
	runSample(t, s, 2)
}
