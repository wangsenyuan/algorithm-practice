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
	s := `())(())(())(
7
1 1
2 3
1 2
1 12
8 12
5 11
2 10
0
0
2
10
4
6
6
`
	runSample(t, s)
}
