package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect answer) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %+v, but got %+v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 3
3 4 5
3 9 1
`, answer{cost: 392, li: 1, lj: 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
1 0 0 0
0 0 3 0
0 0 5 5
`, answer{cost: 240, li: 2, lj: 3})
}
