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
	s := "1 3"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "2 4"
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "199999 200000"
	expect := 36
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "19 84"
	expect := 263
	runSample(t, s, expect)
}
