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
	s := `3 3 3
2 2 2
1 1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 7 9
5 2 7 3
3 5 2 7 3 8 7
`
	expect := false
	runSample(t, s, expect)
}
