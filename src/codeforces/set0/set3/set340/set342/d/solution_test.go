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
	runSample(t, `5
....X
.O...
...X.
1
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
.....
.O...
.....
2
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
...
...
..O
4
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `8
.X.....X
....O...
X.......
0
	`)
}
