package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, c, res := drive(reader)
	expect := bruteForce(k, c)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
1 1 4 2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 100
2 2 2
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10 20
6 4 7 10 4 5 5 3 7 10
`
	runSample(t, s)
}
