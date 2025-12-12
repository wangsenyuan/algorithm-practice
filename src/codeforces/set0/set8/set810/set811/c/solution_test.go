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
4 4 2 5 2 3`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9
5 1 3 1 5 2 4 2 5
`
	expect := 9
	runSample(t, s, expect)
}
