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
	s := `4 4
3 4 3 6
1 2
1 3
2 4
3 4
`
	runSample(t, s, 5)
}

func TestSample2(t *testing.T) {
	s := `4 6
1 1 1 2
1 2
2 3
3 1
1 4
2 4
3 4
`
	runSample(t, s, 9)
}

func TestSample3(t *testing.T) {
	s := `8 11
2 4 2 6 8 10 18 26
1 2
2 3
3 1
4 3
2 4
3 5
5 6
4 6
6 7
7 5
5 8
`
	runSample(t, s, 24)
}

func TestSample4(t *testing.T) {
	s := `2 2
10 10
1 2
2 1
`
	runSample(t, s, 2)
}
