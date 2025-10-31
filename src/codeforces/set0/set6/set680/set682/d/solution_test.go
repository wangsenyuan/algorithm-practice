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
	s := `3 2 2
abc
ab
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 12 4
bbaaababb
abbbabbaaaba
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15 9 4
ababaaabbaaaabb
bbaababbb
`
	expect := 8
	runSample(t, s, expect)
}
