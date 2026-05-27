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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
10 20 30
1 1 1 1 50
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
10 20 30
1 1 1 1 51
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
30 20 10
34 19 50 33 88 15 20
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
10 5 10
10 9 5 25 20 5
`
	expect := 3
	runSample(t, s, expect)
}

func TestPairWithLargestSingle(t *testing.T) {
	s := `2
10 4 5
6 10
`
	expect := 1
	runSample(t, s, expect)
}

func TestPreferThreeSinglesOverPair(t *testing.T) {
	s := `7
11 7 8
6 20 2 1 9 11 1
`
	expect := 3
	runSample(t, s, expect)
}
