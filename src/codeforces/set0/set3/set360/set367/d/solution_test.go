package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2 2
1 2
2 1 3
`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `5 1 1
5 4 5 3 2 1
`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `7 3 1
4 1 3 5 7
2 2 6
1 4
`
	runSample(t, s, 3)
}
