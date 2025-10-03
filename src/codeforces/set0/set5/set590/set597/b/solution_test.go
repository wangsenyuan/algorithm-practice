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
	s := `2
7 11
4 7
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
2 3
3 4
4 5
5 6
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
4 8
1 5
4 7
2 5
1 3
6 8
`
	expect := 2
	runSample(t, s, expect)
}
