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
	runSample(t, `1 0 0
0
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 -2 -1
-2
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 -1 -1
-1 1 -1
2`)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 -3 -2
-1 -1 -1 -2 -1 -1
0`)
}

func TestSample5(t *testing.T) {
	runSample(t, `8 3 2
2 2 -1 -2 3 -1 2 2
2`)
}

func TestSample6(t *testing.T) {
	runSample(t, `9 6 3
1 2 3 1 2 3 1 2 3
7`)
}

func TestSample7(t *testing.T) {
	runSample(t, `13 7 3
0 -1 3 3 3 -2 1 2 2 3 -1 0 3
8`)
}
