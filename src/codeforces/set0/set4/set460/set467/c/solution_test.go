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
	s := `5 2 1
1 2 3 4 5
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 1 3
2 10 7 18 5 33 0
`
	expect := 61
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `13 8 1
73 7 47 91 54 74 99 11 67 35 84 18 19
`
	expect := 515
	runSample(t, s, expect)
}
