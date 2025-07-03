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

		if x != expect {
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
1 1 1
1 1
2 2
3 2
1 2
6
4
7
12
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `20 13
2 2 0 4 3 3 0 2 0 3 4 1 4 3 0 2 4 3 2 2
11 1
6 0
1 2
13 0
10 4
13 1
8 0
6 0
9 4
5 1
19 4
11 2
15 4
55
49
49
45
44
50
48
48
56
54
50
48
52
56
58
54
47
47
`)
}
