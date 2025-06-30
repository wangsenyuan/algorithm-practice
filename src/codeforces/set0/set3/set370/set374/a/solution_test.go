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
	s := `5 7 1 3 2 2
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 5 2 3 1 1
-1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 5 1 3 1 1
-1`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `1 5 1 3 10 1
-1`
	runSample(t, s)
}
