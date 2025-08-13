package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 2 17
3 1 2 5
4 2
3 4
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `3 2 6
3 1 1
1 2
2 3
`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `3 2 10
1 2 3
1 2
2 1
`
	runSample(t, s, 0)
}
