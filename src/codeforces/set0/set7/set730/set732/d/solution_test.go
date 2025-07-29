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
	runSample(t, `7 2
0 1 0 2 1 0 2
2 1
5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 3
0 0 1 2 3 0 2 0 1 2
1 1 4
9`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 1
1 1 1 1 1
5
-1`)
}
