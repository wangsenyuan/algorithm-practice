package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 3 2
1 2
2 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 3 9
1 3
6 1
4 1
1 2
2 1
7 1
9 3
8 1
9 2
`
	expect := 43
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10 36
3 8
5 10
3 10
6 10
2 10
2 8
7 10
1 10
1 8
7 6
7 8
2 5
1 6
8 8
7 5
2 4
9 8
7 4
4 3
10 10
10 8
8 10
10 6
6 2
4 2
10 5
8 3
1 2
2 1
4 1
10 4
10 3
8 1
6 1
10 2
9 1
`
	expect := 153
	runSample(t, s, expect)
}
