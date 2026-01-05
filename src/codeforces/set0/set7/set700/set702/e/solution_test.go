package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	for _, cur := range res {
		var x, y int
		fmt.Fscan(reader, &x, &y)

		if cur[0] != x || cur[1] != y {
			t.Fatalf("Sample expect (%d %d), but got (%d %d)", x, y, cur[0], cur[1])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 3
1 2 3 4 3 2 6
6 3 1 4 2 2 3
10 1
8 1
7 1
10 2
8 2
7 1
9 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 4
0 1 2 3
0 1 2 3
0 0
4 1
8 2
12 3
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 3
1 2 3 4 0
4 1 2 14 3
7 1
17 1
19 2
21 3
8 1
`
	runSample(t, s)
}
