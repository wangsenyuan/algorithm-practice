package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
2 1 3
1 3
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 5
1 2 3 4 5 6 7 8 9 10
2 4 6 8 10
30`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `14 6
7 6 10 9 11 8 14 3 1 13 12 4 5 2
7 10 11 12 4 5
64`
	runSample(t, s)
}
