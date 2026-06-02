package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	tc := len(expect)
	for i := range tc {
		res := drive(reader)
		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2
01
10
1000010011
12
111100010011
3
111
8
00010101`
	expect := []int{0, 2, 3, 0, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
10011
10
1111111111
7
0000000
`
	expect := []int{1, 0, 0}
	runSample(t, s, expect)
}
