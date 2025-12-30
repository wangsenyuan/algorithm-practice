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
	s := `3 2
5 9 3`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4
12 14
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 3
1 1
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20 100
52 55 53 54 48 47 50 45 52 51 51 51 47 45 50 51 53 55 51 52
`
	expect := 6
	runSample(t, s, expect)
}
