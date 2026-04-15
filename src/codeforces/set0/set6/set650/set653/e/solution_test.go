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
	s := `5 4 2
1 2
2 3
4 2
4 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 5 3
1 2
1 3
1 4
1 5
1 6
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3 2
2 3
2 4
3 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 2 2
1 2
1 3
`
	expect := false
	runSample(t, s, expect)
}

