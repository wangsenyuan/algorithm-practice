package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
2 3
3 2
3 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 2
2 1
2 2
`
	expect := 2
	runSample(t, s, expect)
}
