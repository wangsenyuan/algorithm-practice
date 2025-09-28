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
	s := `9 9 1
1 2
1 3
2 3
1 5
5 6
6 1
1 8
9 8
7 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4 5
1 2
2 3
3 4
4 1
`
	expect := 1
	runSample(t, s, expect)
}
