package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2 3 4 5
3 2 7 1 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
0 0 1
1 0 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
0 0 1
0 0 0
`
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := `4
0 0 1 2
1 3 3 2
`
	runSample(t, s, false)
}

func TestSample5(t *testing.T) {
	s := `6
1 1 4 5 1 4
0 5 4 5 5 4
`
	runSample(t, s, true)
}

func TestSample6(t *testing.T) {
	s := `3
0 1 2
2 3 2
`
	runSample(t, s, false)
}

func TestSample7(t *testing.T) {
	s := `2
10 10
11 10
`
	runSample(t, s, false)
}
