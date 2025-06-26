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
	s := `4
1 2 1 2
4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
2 2
1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10
1 2 1 5 1 2 2 1 1 2
16`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `8
1 5 2 8 4 1 4 2
16`
	runSample(t, s)
}

