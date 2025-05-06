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
			t.Fatalf("Sample expect %s, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 4 1 4 1
2 5 2 6 2
4
0 0
3 3
4 5
3 5
`
	expect := `0
1
2
0
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
0 3 2 7 2
0 5 2 10 12
2 5 2 9 30
1 8 1 8 11
2 4 2 7 17
5
8 3
0 8
8 0
8 2
0 2
`
	expect := `4
2
0
0
1
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
2 7 2 8 12
0 4 0 5 30
2 6 1 6 32
1 5 0 6 35
1 5 2 3 26
0 1 1 6 48
1 7 0 2 73
0 2 1 6 41
5
0 6
4 5
1 4
0 4
1 4
`
	expect := `8
1
2
6
4
`
	runSample(t, s, expect)
}
