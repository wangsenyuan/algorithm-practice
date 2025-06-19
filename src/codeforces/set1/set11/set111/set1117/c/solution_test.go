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
	runSample(t, `0 0
4 6
3
UUU
5
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `0 3
0 0
3
UDD
3
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `0 0
0 1
1
L
-1
	`)
}
