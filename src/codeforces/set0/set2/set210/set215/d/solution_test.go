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
	s := `2 10
30 35 1 100
20 35 10 10
120`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 100
10 30 1000 1
5 10 1000 3
10 40 1000 100000
200065`
	runSample(t, s)
}
