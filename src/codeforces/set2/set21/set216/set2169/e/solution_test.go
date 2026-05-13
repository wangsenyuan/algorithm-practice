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
	s := `1
42
42
1000
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
5 10 5 0
0 5 10 5
1 1 1 1
`
	expect := 40
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
6 7 8 9
3 3 3 3
9 0 9 0
`
	expect := 22
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
1000000000 10
10 1000000000
12345 54321
`
	expect := 3999999960
	runSample(t, s, expect)
}

func TestPositiveCoordinatesNeedNegativePartialState(t *testing.T) {
	s := `2
10 20
10 20
0 0
`
	expect := 40
	runSample(t, s, expect)
}
