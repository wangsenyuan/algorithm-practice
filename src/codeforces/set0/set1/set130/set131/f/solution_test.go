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
	runSample(t, `4 6 2
111000
111100
011011
000111
6`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 5 4
11111
11111
11111
11111
11111
9`)
}