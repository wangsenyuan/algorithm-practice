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
	runSample(t, `3
1 2 3
1 2
2 3
10`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
1 2
2 3
3 4
3 5
52`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
10 9 8 7 6
1 2
2 3
3 4
3 5
131`)
}
