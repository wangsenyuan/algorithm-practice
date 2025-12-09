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
	s := `4 2
9 9
10 8
5 3
4 3
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4
1 2 3 4
10 3 7 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 1
3
6
2
5
1
4
`
	expect := 3
	runSample(t, s, expect)
}
