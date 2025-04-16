package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
0 -2 3 -5 1
2
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
1 -3 -10 4 1
3
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
-2 -5 4
1
`
	expect := 11
	runSample(t, s, expect)
}
