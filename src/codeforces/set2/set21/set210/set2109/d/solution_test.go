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
	runSample(t, `6 5 2
2 3
1 2
2 3
3 4
4 5
5 6
111101`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 5 1
5
1 2
2 3
3 4
4 5
3 5
11111`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 4 3
100 200 300
1 2
1 3
1 4
2 5
10001`)
}
