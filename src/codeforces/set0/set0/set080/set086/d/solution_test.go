package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	reader = bufio.NewReader(strings.NewReader(expect))

	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2 1
1 2
1 3
`
	expect := `3
6
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 3
1 1 2 2 1 3 1 1
2 7
1 6
2 7
`
	expect := `20
20
20
`
	runSample(t, s, expect)
}
