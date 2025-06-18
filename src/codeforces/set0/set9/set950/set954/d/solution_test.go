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
	runSample(t, `5 4 1 5
1 2
2 3
3 4
4 5
0
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4 3 5
1 2
2 3
3 4
4 5
5
`)
}


func TestSample3(t *testing.T) {
	runSample(t, `5 6 1 5
1 2
1 3
1 4
4 5
3 5
2 5
3
`)
}

