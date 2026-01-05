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
	s := `7
2 1 5 11 5 9 11
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
5 4 3 2 1
`
	expect := 12
	runSample(t, s, expect)
}
