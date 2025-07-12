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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `13
12
12`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `16
11311
475`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `20
999
3789`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `17
2016
594`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `1000
1001
100001`
	runSample(t, s)
}
