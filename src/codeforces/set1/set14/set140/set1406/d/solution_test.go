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
	s := `4
2 -1 7 3
2
2 4 -3
3 4 2
5
5
6
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6
-9 -10 -9 -6 -5 4
3
2 6 -9
1 2 -10
4 6 -3
3
3
3
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1
0
2
1 1 -1
1 1 -1
0
0
-1
`
	runSample(t, s)
}
