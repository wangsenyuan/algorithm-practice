package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 4 2
4 3 2
2 5 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 8 10 0 7
7 7 3 0 10
2 8 3 2 2
`
	expect := 1
	runSample(t, s, expect)
}
