package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
0 2
1 0
2 1
2 3
`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `5
0 1
0 2
0 3
0 4
0 5
`
	runSample(t, s, 5)
}

func TestSample3(t *testing.T) {
	s := `8
10 4
4 2
7 10
5 3
1 9
0 6
3 8
0 9
`
	runSample(t, s, 2)
}
