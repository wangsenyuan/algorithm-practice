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
	s := `5 1 10
01000
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 10 1
01000
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 2 3
1111111
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `14 3 11
10110100011001
`
// 101010101
// 10101
	expect := 20
	runSample(t, s, expect)
}