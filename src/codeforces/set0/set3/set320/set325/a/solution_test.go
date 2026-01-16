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
	s := `5
0 0 2 3
0 3 3 5
2 0 5 2
3 2 5 5
2 2 3 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
0 0 2 3
0 3 3 5
2 0 5 2
3 2 5 5
`
	expect := false
	runSample(t, s, expect)
}