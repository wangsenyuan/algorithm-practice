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
	runSample(t, `3 3
1 2 4
2 3 5
1 3 2
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4
1 4 7
4 2 2
2 3 4
3 4 1
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `999 4
1 2 9
2 1 8
1 2 7
1 1 6
-1`)
}
