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
	s := `13 5
2 4 5 5 4 3 2 2 2 3 3 2 1
3 4 4 3 2
`
	expect := 2
	runSample(t, s, expect)
}
