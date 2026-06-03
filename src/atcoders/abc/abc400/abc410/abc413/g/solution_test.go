package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1 0`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 12 20
8 3
1 11
6 4
3 7
10 4
5 7
4 7
5 5
4 3
6 1
1 6
2 7
6 7
1 3
6 3
2 12
9 6
7 3
3 11
9 7
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 7 3
1 2
2 4
1 6
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 5 5
1 4
2 3
3 2
3 4
4 2
`
	expect := false
	runSample(t, s, expect)
}
