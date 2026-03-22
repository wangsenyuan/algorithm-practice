package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
0111
0001
`
	expect := "First"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
110110
001001
`
	expect := "First"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
111000
000111
`
	expect := "Draw"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
01010110
00101101
`
	expect := "First"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
01100000
10010011
`
	expect := "Second"
	runSample(t, s, expect)
}
