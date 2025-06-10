package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		expect := readNum(reader)
		if expect != x {
			t.Fatalf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
2 3 1 5 5
4
2 2 3 2
2 1 5 3
1 3 5
2 1 5 3
5
26
38
`
	runSample(t, s)
}
