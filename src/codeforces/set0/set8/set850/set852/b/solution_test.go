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
	s := `2 3 13
4 6
2 1
3 4
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4 5
1 1
1 1
1 1
`
	expect := 16
	runSample(t, s, expect)
}
