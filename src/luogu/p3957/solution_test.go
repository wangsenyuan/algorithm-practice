package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 4 10
2 6
5 -3
10 3
11 -3
13 1
17 6
20 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 4 20
2 6
5 -3
10 3
11 -3
13 1
17 6
20 2`
	expect := -1
	runSample(t, s, expect)
}
