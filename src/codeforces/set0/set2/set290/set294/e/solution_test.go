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
	s := `3
1 2 2
1 3 4
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 1
2 3 1
3 4 1
4 5 1
5 6 1
`
	expect := 29
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 3 1
2 3 1
3 4 100
4 5 2
4 6 1
`
	expect := 825
	runSample(t, s, expect)
}
