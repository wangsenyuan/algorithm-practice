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
1 2 3
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
2 3 4 5
`
	expect := 19
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 4 4 4 4 4
`

	expect := 38
	runSample(t, s, expect)
}
