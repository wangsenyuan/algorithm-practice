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
	runSample(t, `3 2
1 2 1
2 3 1
2
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 2
1 3 13
1 2 5
18
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 9
1 2 6
2 3 5
3 8 6
1 4 7
4 5 4
5 8 7
1 6 5
6 7 5
7 8 5
10
`)
}

func TestSample4(t *testing.T) {
	runSample(t, `3 3
1 3 9
1 2 8
2 3 3
11
`)
}
