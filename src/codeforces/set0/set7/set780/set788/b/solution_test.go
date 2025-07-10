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
	runSample(t, `5 4
1 2
1 3
1 4
1 5
6`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
1 1
1 2
1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 4
2 3
2 4
3 4
4 4
6`)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 3
1 2
2 3
4 4
0`)
}
