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
	s := `3 2
2 1 2
1 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 3
3 1 3 7
2 2 5
2 4 6
`
	expect := 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 5
2 1 2
2 3 4
1 5
2 6 7
1 8
`
	expect := 8
	runSample(t, s, expect)
}
