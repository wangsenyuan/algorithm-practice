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
1 3 2
`

	//  (1, 2), (1, 3), (2, 3)

	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
1 3 2 1 7
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 3
1 7 6 4 9 5 3
`
	// 0 + 4 + 3 + 1 + 2 + 1 = 11
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `40 1000000000000000000
83 35 47 18 96 63 24 91 15 100 40 23 20 34 65 22 52 87 55 19 11 73 45 28 60 61 24 42 30 43 65 75 31 84 100 12 69 98 49 25
`
	expect := 780
	runSample(t, s, expect)
}
