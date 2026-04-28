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
		t.Errorf("expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 4 4 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 3 3`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 6 9 2`
	expect := 6
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 6 9 3`
	expect := 8
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 6 7 4`
	expect := -1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 5 5 1`
	expect := -1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `2 3 6 2`
	expect := -1
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `999999999999999999 1000000000000000000 1000000000000000000 999999999999999999`
	expect := 1000000000000000000
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `1000000000000000000 1 999999999999999999 1000000000000000000`
	expect := 2
	runSample(t, s, expect)
}
