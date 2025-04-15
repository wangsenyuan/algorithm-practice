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
	s := `3 1000
1000
0 0
0 1
0 3
`
	expect := 2000
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1000
1000
1 0
1 1
1 2
`
	expect := 1000
	runSample(t, s, expect)
}
