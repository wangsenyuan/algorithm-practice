package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
5 4 3 2 1 2 3 4 5
55555
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
9 11 1 12 5 8 9 10 6
33
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `0
1 1 1 1 1 1 1 1 1
-1
	`)
}
