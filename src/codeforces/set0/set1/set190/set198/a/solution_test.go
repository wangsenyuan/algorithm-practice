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
	runSample(t, `3 1 3 5`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 4 4 7`, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 2 4 100`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 1 1000000 1000000`, 1)
}