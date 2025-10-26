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
	s := `4
1 2 2 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 3 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
2 3 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `16
1 4 13 9 11 16 14 6 5 12 7 8 15 2 3 10
`
	expect := 105
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1
1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `20
11 14 2 10 17 5 9 6 18 3 17 7 4 15 17 1 4 14 10 11
`
	expect := 7
	runSample(t, s, expect)
}
