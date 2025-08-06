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
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `1
42
`
	expect := 42
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 2
`
	expect := 13
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2 3 4 5
`
	expect := 719
	runSample(t, s, expect)
}
