package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readString(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 3
5
9 1 1000000000
YES`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 4 3
3 4
NO`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4 3
2 4 6 5
6 1 8
YES`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `5 2
6 4 5 4 5
4 1000
NO`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `3 1
9 8 7
8
YES`
	runSample(t, s)
}
