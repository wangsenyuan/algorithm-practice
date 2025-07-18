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
3 7 8
1
1 3
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `7
1 2 1 3 3 2 3
5
4 7
4 5
1 3
1 7
1 5
0
3
1
3
2`
	runSample(t, s)
}
