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
	runSample(t, `2
1 2
4`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1 1 2 3
20`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 1 1 1 1
62`)
}
