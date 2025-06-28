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
			t.Fatalf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
5 10 2 1
? 1 2 1
= 2 2 0
? 2 4 3
= 1 4 1
? 1 4 5
25
43
1300
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 1
1000000000 1000000000 1000000000
? 1 3 0
999999986
`
	runSample(t, s)
}
