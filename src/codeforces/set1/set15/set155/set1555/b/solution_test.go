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
		t.Errorf("Sample %q expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8 5
2 1 7 4
4 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
2 2 5 4
3 3`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 8
0 3 1 6
1 5`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 1
3 0 6 1
5 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 10
4 5 7 8
8 5`
	expect := 0
	runSample(t, s, expect)
}
