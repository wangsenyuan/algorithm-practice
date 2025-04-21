package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
3 6 4 1 2`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 1
3 1 4 1 5 9 2`
	expect := 16
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3
4 3 2 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 2
1 3 5 2 4 6`
	expect := 6
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 1
1000000000 1 1000000000 1 1000000000 1`
	expect := 3000000000
	runSample(t, s, expect)
}
