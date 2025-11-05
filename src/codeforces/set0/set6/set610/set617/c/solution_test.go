package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 -1 0 5 3
0 2
5 2`, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 0 0 5 0
9 4
8 3
-1 0
1 4
`, 33)
}
