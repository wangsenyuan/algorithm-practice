package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
...
..#
...`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2
..
..
..
..`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 5
#...#
#...#
###.#
###.#`, 4)
}
