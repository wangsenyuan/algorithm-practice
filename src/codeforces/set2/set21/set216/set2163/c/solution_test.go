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
	s := `2
1 3
3 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 3
3 2 1
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 5 5 5
5 3 1 2
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
8 8 8 8
8 8 8 8
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
6 6 5 7 9 12
1 4 2 8 5 6
`
	expect := 25
	runSample(t, s, expect)
}

