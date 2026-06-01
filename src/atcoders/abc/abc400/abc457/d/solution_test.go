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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `3 3
1 2 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
10 1 10 1
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 457
8 9 10 9 8 8 4 6 8 1 5 10 2 8 2 6 8 1 6 6
`
	expect := 132
	runSample(t, s, expect)
}
