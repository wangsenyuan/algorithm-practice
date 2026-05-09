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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 4 3 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 3 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
2 1`
	expect := 1
	runSample(t, s, expect)
}