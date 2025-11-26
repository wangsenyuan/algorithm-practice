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
	s := `3
1 3 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 2 3 4 5
`
	expect := true
	runSample(t, s, expect)
}