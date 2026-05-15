package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1 5
3
`
	runSample(t, s, 0)
}

func TestSample2(t *testing.T) {
	s := `2 100 100
50 200
`
	runSample(t, s, 150)
}

func TestSample3(t *testing.T) {
	s := `5 1 10
5 7 3 9 1
`
	runSample(t, s, 12)
}

func TestSample4(t *testing.T) {
	s := `5 6 10
9 3 1 7 5
`
	runSample(t, s, 13)
}
