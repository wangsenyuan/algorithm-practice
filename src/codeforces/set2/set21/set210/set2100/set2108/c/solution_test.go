package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5
4 3 2 1 5
2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 1 1
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6
7 8 1 5 9 2
2
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `10
1 7 9 7 1 10 2 10 10 7
3
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `3
1 1 2
1
`
	runSample(t, s)
}