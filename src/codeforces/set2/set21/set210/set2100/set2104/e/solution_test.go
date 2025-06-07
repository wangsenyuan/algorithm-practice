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
	s := `7 3
abacaba
3
cc
bcb
b
0
1
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 1
aaaaa
6
a
aa
aaa
aaaa
aaaaa
aaaaaa
5
4
3
2
1
0
`
	runSample(t, s)
}
