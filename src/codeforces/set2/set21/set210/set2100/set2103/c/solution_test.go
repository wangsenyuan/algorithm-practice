package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)

	if expect == "YES" != res {
		t.Errorf("Sample expect %s, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2
3 2 1
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
3 2 1
NO`)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 3
8 5 3 1 6 4
NO`)
}

func TestSample4(t *testing.T) {
	runSample(t, `8 7
10 7 12 16 3 15 6 11
YES`)
}

func TestSample5(t *testing.T) {
	runSample(t, `6 8
7 11 12 4 9 17
YES`)
}

func TestSample6(t *testing.T) {
	runSample(t, `3 500000000
1000 1000000000 1000
YES`)
}
