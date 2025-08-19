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
	s := `8
1 2 7 3 4 8 5 6
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
25 1 2 3 14 36
`
	expect := 2
	runSample(t, s, expect)
}
