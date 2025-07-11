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
	runSample(t, `6 2
1 2
1 3
2 4
2 5
4 6
20`)
}

func TestSample2(t *testing.T) {
	runSample(t, `13 3
1 2
3 2
4 2
5 2
3 6
10 6
6 7
6 13
5 8
5 9
9 11
11 12
114`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 5
2 1
3 1
3`)
}
