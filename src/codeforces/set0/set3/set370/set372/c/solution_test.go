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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `50 3 1
49 1 1
26 1 4
6 1 10
-31`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 2 1
1 1000 4
9 1000 4
1992`
	runSample(t, s)
}
