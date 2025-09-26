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
	s := `4 5 2
0 2 1 1 0
0 0 3 1 2
1 0 4 3 1
3 3 0 0 4
`
	expect := 25
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3 1
1 2 3
4 5 6
7 8 9
`
	expect := 31
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 2
1 2 3
4 5 6
7 8 9
`
	expect := 44
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 3 3
1 2 3
4 5 6
7 8 9
`
	expect := 45
	runSample(t, s, expect)
}
