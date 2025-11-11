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
	s := `6
3 5 1 4 3 2
LRLLLR
`
	expect := 18
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 8
LR
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2 3 4 5
LRLRR
`
	expect := 22
	runSample(t, s, expect)
}
