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
	s := `2 12 0 4
3 4
4 2
`
	runSample(t, s, 6)
}

func TestSample2(t *testing.T) {
	s := `4 1000 10 35
1 2 4 5
1000 500 250 200
`
	runSample(t, s, 5)
}

func TestSample3(t *testing.T) {
	s := `2 10 7 11
2 10
6 1
`
	runSample(t, s, 11)
}
