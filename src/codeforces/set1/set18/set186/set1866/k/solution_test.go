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
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7
5 1 2
1 4 2
3 4 1
2 5 3
6 1 6
4 7 2
2
4 3
3 2
`
	expect := `18
11
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 1000000000
2 3 1000000000
1
2 1000000000
`
	expect := `2000000000000000000`
	runSample(t, s, expect)
}
