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
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
1 2
2 3
3 2
3 1
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 6
1 2
2 3
3 2
3 1
2 1
4 5
NO`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2
1 2
2 1
YES`)
}
