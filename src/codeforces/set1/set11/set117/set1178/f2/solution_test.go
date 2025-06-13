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
1 2 3
5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `7 7
4 5 1 6 2 3 7
165`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 3
1 2 1
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 3
2 1 2
0`)
}

func TestSample5(t *testing.T) {
	runSample(t, `8 17
1 3 2 2 7 8 2 5 5 4 4 4 1 1 6 1 1
20`)
}
