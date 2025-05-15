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
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
2 4 1
3
6
8
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 2
-1
-1
`
	runSample(t, s)
}

