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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
24 21 14 10
3
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
14 5 1
-1
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 3 6 9 12
1
`)
}
