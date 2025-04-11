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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
4 3 2 5
`
	expect := 39
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 1 2 7 1
`
	expect := 49
	runSample(t, s, expect)
}
