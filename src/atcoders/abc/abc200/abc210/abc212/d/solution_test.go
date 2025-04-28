package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)

	reader = bufio.NewReader(strings.NewReader(expect))

	for _, x := range ans {
		y := readNum(reader)
		if y != x {
			t.Errorf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 3
1 5
3
2 2
3
`
	expect := `3
7
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 1000000000
2 1000000000
2 1000000000
2 1000000000
2 1000000000
3
`
	expect := `5000000000`
	runSample(t, s, expect)
}
