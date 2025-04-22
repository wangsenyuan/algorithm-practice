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

	m := readNum(reader)

	if len(res) != m {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	for _, cur := range res {
		c, v := readTwoNums(reader)
		if c != cur[0] || v != cur[1] {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 4
1 2 5
2 4 0
1 3 3
1 3 2
`
	expect := `3
0 5
2 4
5 3
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1 5
1 1 1
1 1 10
2 1 100
1 1 1000
2 1 10000
`
	expect := `1
10000 1
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5 10
1 1 1
1 2 2
1 3 3
1 4 4
1 5 5
2 1 6
2 2 7
2 3 8
2 4 9
2 5 10
`
	expect := `5
6 5
7 5
8 5
9 5
10 5
`
	runSample(t, s, expect)
}
