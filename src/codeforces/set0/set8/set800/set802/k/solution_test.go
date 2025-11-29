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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `9 3
0 1 1
0 2 1
1 3 2
1 4 2
1 5 2
2 6 3
2 7 3
2 8 3
`
	expect := 15
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9 5
0 1 1
0 2 1
1 3 2
1 4 2
1 5 2
2 6 3
2 7 3
2 8 3
`
	expect := 17
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `11 6
1 0 7932
2 1 1952
3 2 2227
4 0 9112
5 4 6067
6 0 6786
7 6 3883
8 4 7137
9 1 2796
10 5 6200
`
	expect := 54092
	runSample(t, s, expect)
}
