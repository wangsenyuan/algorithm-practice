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
	s := `10 2 1
1 2
1 3
3 4
3 5
2 6
6 8
6 7
9 8
8 10
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1000 1000
1 2
1 3
3 4
3 5
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3 5
1 2
2 3
3 4
4 5
5 6
`

	// 0, 1, 1, 1
	// 0  2, 1, 0
	// 0, 3, 0, 0

	expect := 3
	runSample(t, s, expect)
}
