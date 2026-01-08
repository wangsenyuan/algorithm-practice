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
	s := `5 2 1 4 10
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2 1 4 5
`
	expect := 13
	runSample(t, s, expect)
}
