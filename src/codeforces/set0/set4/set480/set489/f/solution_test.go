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
	s := `3 1 1000
011
`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `4 4 100500
0110
1010
0101
1001
`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `5 0 13`
	runSample(t, s, 12)
}
