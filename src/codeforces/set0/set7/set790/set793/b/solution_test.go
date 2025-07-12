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
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
..S..
****.
T....
****.
.....
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 5
S....
****.
.....
.****
..T..
NO`)
}
