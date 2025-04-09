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
		if x != y {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `caaaba
5
1 1
1 4
2 3
4 6
4 5
`
	expect := `1
7
3
4
2
`
	runSample(t, s, expect)
}
