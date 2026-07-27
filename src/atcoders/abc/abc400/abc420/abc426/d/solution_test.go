package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if len(res) != 1 || res[0] != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
01001
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3
000
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
15
110010111100101
`, 16)
}
