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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 2
1 1 1
1 2 3 4 5 6
`
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `0 0 10
3 2 3
1 2 3 4 5 6
`
	expect := 4
	runSample(t, s, expect)
}
