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
	s := `3 1
1 1 2 2
2 2 3 3
3 3 4 4
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1
1 1 2 2
1 9 2 10
9 9 10 10
9 1 10 2
`
	expect := 64
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 0
1 1 2 2
1 1 1000000000 1000000000
1 3 8 12
`
	expect := 249999999000000001
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `11 8
9 1 11 5
2 2 8 12
3 8 23 10
2 1 10 5
7 1 19 5
1 8 3 10
1 5 3 9
1 2 3 4
1 2 3 4
4 2 12 16
8 5 12 9
`
	expect := 4
	runSample(t, s, expect)
}
