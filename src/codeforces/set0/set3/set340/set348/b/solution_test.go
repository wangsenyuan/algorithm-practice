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
	runSample(t, `6
0 0 12 13 5 6
1 2
1 3
1 4
2 5
2 6
6`)
}

func TestSample2(t *testing.T) {
	runSample(t, `10
0 9 5 0 0 0 0 0 9 7
7 5
8 1
1 5
4 3
2 4
4 7
7 9
10 6
6 8
22`)
}
