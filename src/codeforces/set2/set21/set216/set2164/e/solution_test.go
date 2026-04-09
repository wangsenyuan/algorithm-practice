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
	s := `5 6
2 4 15
2 5 4
1 3 6
2 3 9
1 2 10
3 4 7
`
	expect := 58
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
1 2 3
1 3 2
1 4 1
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 3
1 2 1
2 1 3
1 1 4
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 6
2 3 10
1 3 10
5 6 10
6 6 1
4 5 10
3 4 10
`
	expect := 71
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 5
1 2 4
5 1 5
4 3 6
2 4 10
1 4 7
`
	expect := 43
	runSample(t, s, expect)
}
