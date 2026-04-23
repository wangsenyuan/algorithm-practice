package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2 4
3 4 2 3 1
3 2 3 4 1
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `3 1 2
1 2 3
3 1 2
`
	runSample(t, s, false)
}

func TestSample3(t *testing.T) {
	s := `4 2 4
1 1 1 1
1 1 1 1
`
	runSample(t, s, true)
}
