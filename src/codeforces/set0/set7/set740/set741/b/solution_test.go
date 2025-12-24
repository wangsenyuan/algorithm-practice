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
	s := `3 1 5
3 2 5
2 4 2
1 2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2 11
2 4 6 6
6 4 2 1
1 2
2 3
`
	expect := 7
	runSample(t, s, expect)
}
