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
	s := `5
2 2 1 2 1
`
	runSample(t, s, 12)
}

func TestSample2(t *testing.T) {
	s := `4
4 4 4 4
`
	runSample(t, s, 4)
}

func TestSample3(t *testing.T) {
	s := `10
1 2 1 4 3 3 3 2 2 4
`
	runSample(t, s, 47)
}
