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
	s := `0 1 1 2
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `0 2 0 2
1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 3 1 3
4`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `9 98 244 353
374`
	runSample(t, s)
}