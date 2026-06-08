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
	s := `3
1 1
`
	runSample(t, s, 4)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 2 1
`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `15
4 4 4 4 4 4 3 2 2 2 2 2 1 1
`
	runSample(t, s, 70270200)
}

func TestSample4(t *testing.T) {
	s := `2
1
`
	// [2,1], [1, 2]
	runSample(t, s, 2)
}
