package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		expect := readNum(reader)
		if x != expect {
			t.Fatalf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 3
2 4 2 1 3 2
7
4
6
2
4
3`)
}
