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
	s := `3
0 1 2 1
1 4 1 2
0 3 2 3
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
-2 -1 2 -1
2 1 -2 1
-1 -2 -1 2
1 2 1 -2
`
	expect := 16
	runSample(t, s, expect)
}
