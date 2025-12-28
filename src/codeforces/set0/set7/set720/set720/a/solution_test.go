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
	s := `2 2
3 3 3 2
1 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
3 2 3 3
1 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1
0
1 1
`
	expect := false
	runSample(t, s, expect)
}
