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
	runSample(t, `3 0`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4
1 2 1
2 3 1
3 4 0
4 1 0
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 4
1 2 1
2 3 1
3 4 0
4 1 1
`, 0)
}
