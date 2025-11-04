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
	s := `5 3
3-a 2-b 4-c 3-a 2-c
2-a 2-b 1-c
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
3-a 6-b 7-a 4-c 8-e 2-a
3-a
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5
1-h 1-e 1-l 1-l 1-o
1-w 1-o 1-r 1-l 1-d
`
	expect := 0
	runSample(t, s, expect)
}
