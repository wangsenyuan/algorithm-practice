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
	runSample(t, `3
2 1 3
7
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
2 1 1 3
11
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1 5 2 5 5 3 10 6 5 1
104
	`)
}
