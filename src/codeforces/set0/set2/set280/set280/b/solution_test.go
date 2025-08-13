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
	s := `5
5 2 1 4 3
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
9 8 3 5 7
`
	expect := 15
	runSample(t, s, expect)
}
