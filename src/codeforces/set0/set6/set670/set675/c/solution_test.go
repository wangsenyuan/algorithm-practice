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
	s := `3
5 0 -5
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
-1 0 1 0
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 3 -6
`
	// 0 1 2 3 0 1 2 3 0
	expect := 3
	runSample(t, s, expect)
}
