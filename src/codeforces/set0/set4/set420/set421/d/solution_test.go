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
	s := `4 2
2 3
1 4
1 4
2 1
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 6
5 6
5 7
5 8
6 2
2 1
7 3
1 3
1 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3
4 5
4 5
4 5
3 2
2 3
`
	expect := 7
	runSample(t, s, expect)
}
