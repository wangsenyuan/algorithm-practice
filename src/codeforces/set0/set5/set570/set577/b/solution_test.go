package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5
1 2 3`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 6
5`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 6
3 1 1 3`, true)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 6
5 5 5 5 5 5
`, true)
}
