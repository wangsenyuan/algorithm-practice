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
	runSample(t, `5 2
..XXX
XX.XX
X.XXX
X...X
XXXX.
10`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3
.....
.XXX.
.XXX.
.XXX.
.....
25`)
}
