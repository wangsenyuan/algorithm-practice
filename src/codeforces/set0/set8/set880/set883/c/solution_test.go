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
	s := `120 964 20
26 8 8
13 10 4`
	expect := 40
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 200 20
1 1 1
2 2 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 81 11
4 10 16
3 10 12
`
	expect := 28
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 79 11
4 10 16
3 10 12
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 12 14
19 2 4
8 1 10
`
	expect := 14
	runSample(t, s, expect)
}
