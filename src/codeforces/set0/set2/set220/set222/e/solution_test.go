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
	s := `3 3 2
ab
ba
17`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 3 0
27`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 1 1
aa
0`
	runSample(t, s)
}
