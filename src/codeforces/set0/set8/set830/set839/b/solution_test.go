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
	s := `2 2
5 8
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
7 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2
4 4
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 4
2 2 1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2 7
2 2 2 2 2 2 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 3
2 2 1
`
	expect := true
	runSample(t, s, expect)
}
