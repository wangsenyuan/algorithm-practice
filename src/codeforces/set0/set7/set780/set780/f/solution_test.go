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
	s := `2 2
1 2 0
2 2 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 2 0
2 2 1
2 2 0
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1
1 1 0
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 6
3 1 1
1 2 1
2 1 1
1 3 0
3 2 1
2 2 0
`
	expect := 30
	runSample(t, s, expect)
}
