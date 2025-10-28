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
	s := `3 5
10000
10010
00001
`
	expect := 20
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
00
00
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
11
11
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 10
0000000000
0000000000
0000001000
0010000000
0100000000
0000000000
0000000000
0000000000
0000000000
0010001000
`
	expect := 1183
	runSample(t, s, expect)
}
