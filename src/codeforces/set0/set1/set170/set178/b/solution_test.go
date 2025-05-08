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
	s := `7 8
1 2
2 3
3 4
4 5
5 6
5 7
3 5
4 7
4
1 5
2 4
2 6
4 7
`
	expect := `2
1
2
0
`
	runSample(t, s, expect)
}
