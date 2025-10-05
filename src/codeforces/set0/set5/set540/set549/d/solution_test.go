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
	s := `6 8
BBBBBBBB
BBBBBBBB
BBBBBBBB
WWWWWWWW
WWWWWWWW
WWWWWWWW
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
WBW
BWW
WWW
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 6
WWBBWW
WWBBWW
WWBBWW
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4
BBBB
BBBB
BBBB
BBBW
`
	expect := 4
	runSample(t, s, expect)
}
