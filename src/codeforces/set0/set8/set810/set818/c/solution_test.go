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
3 2
3 1 3 2
1 2 2 2
1 0 0 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
10 10
1 2 1 1
5 5 6 5
6 4 5 4
2 1 2 0
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2 2
2 1 1 1
1 2 2 2
1 0 0 0
`
	expect := -1
	runSample(t, s, expect)
}
