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
	s := `4 3
1 0 0 1`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
1 1 1 1 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 5
0 1 0 1 0`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 3
1 0 1 0 1 1`
	expect := 16
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 3
1 0 1 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5 3
1 0 1 1 0`
	expect := 7
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `34 17
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1`
	expect := 333606206
	runSample(t, s, expect)
}
