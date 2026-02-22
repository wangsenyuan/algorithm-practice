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
	s := `7
3 4 3 5 7 6 2
`
	runSample(t, s, 4)
}

func TestSample2(t *testing.T) {
	s := `5
5 4 3 2 1
`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `10
1 2 3 4 5 6 7 8 9 10
`
	runSample(t, s, 10)
}
