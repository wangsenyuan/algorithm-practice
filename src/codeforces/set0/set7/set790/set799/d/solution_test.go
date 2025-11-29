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
	s := `3 3 2 4 4
2 5 4 10
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 3 3 5
2 3 5 4 2
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5 1 2 3
2 2 3
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 4 1 1 3
2 3 2
`
	expect := 3
	runSample(t, s, expect)
}
