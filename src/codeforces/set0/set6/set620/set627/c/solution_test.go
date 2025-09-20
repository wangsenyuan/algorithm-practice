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
	runSample(t, `10 4 4
3 5
5 8
6 3
8 4`, 22)
}

func TestSample2(t *testing.T) {
	runSample(t, `16 5 2
8 2
5 1
`, -1)
}
