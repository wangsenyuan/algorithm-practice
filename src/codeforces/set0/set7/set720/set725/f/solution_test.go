package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `2
12 3 4 7
1 15 9 1
`
	runSample(t, s, 1)
}

func TestSample2(t *testing.T) {
	s := `2
5 4 8 8
4 12 14 0
`
	runSample(t, s, 4)
}

func TestSample3(t *testing.T) {
	s := `1
0 10 0 10
`
	runSample(t, s, -10)
}
