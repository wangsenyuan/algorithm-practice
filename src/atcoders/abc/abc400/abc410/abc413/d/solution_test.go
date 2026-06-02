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
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 8 2 4 16
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
-16 24 54 81 -36
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
90000 8100 -27000 729 -300000 -2430 1000000
`
	expect := true
	runSample(t, s, expect)
}
