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
	s := `3 10
4 6 7
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 12
1 10
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 7
3 4
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7 1000000000
1 10001 10011 20011 20021 40021 40031
`
	expect := 999999969
	runSample(t, s, expect)
}
