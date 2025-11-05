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
	s := `3 7 6
10 8 C
4 3 C
5 6 D
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4 5
2 5 C
2 1 D
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 10 10
5 5 C
5 5 C
10 11 D
`
	expect := 10
	runSample(t, s, expect)
}
