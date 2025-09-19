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
	runSample(t, `4 2
1 3 2 4
3 2
2 4`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 5
9 7 2 3 1 4 6 5 8
1 6
4 5
2 7
7 2
2 7`, 20)
}
