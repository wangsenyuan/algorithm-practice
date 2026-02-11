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
	s := `4
1 2 3
1 2 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4
1 1 1 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 1 1 1 2
1 2 1 2 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
1 1 3 4 4 5
1 2 1 4 2 5
`
	expect := 3
	runSample(t, s, expect)
}

