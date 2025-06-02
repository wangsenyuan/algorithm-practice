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
2 0 0
1 2 1
2 3 1
1 3 2
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 6
2 2 5 0 1
1 2 2
1 3 1
1 4 3
3 5 5
2 4 4
4 5 3
4`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 0
1 1
-1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 4
3 10 0 0
1 2 1
1 3 3
2 3 10
3 4 5
10`)
}
