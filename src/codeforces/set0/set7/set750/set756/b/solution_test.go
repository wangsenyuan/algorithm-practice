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
	s := `3
10
20
30
20
20
10
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10
13
45
46
60
103
115
126
150
256
516
20
20
10
0
20
0
0
20
20
10
`
	runSample(t, s)
}
