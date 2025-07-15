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
	runSample(t, `4
3 6 6 3
NO`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
21 18 15 12 9
YES`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
52 101
YES`)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
10 2
NO`)
}
