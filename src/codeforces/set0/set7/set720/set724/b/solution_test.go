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
	runSample(t, `2 4
1 3 2 4
1 3 4 2
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4
1 2 3 4
2 3 4 1
3 4 1 2
4 1 2 3
NO`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 6
2 1 3 4 5 6
1 2 4 3 5 6
1 2 3 4 6 5
YES`)
}
