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
	runSample(t, `3 3
100 100 100
100 1 100
100 100 100
800
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 5
87882 40786 3691 85313 46694
28884 16067 3242 97367 78518
4250 35501 9780 14435 19004
64673 65438 56977 64495 27280
747898
	`)
}
