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
	runSample(t, `3
1 1
1 2
2 1
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 1
2 6
6 4
3 3
8 2
32`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
1 1
1 1000000000
1000000000 1
1000000000 1000000000
1000000000000000000`)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
1 1
1`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
1 2
4 2
4 3
3 1
3 2
6`)
}
