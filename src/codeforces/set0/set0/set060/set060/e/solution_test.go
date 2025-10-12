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
	s := `2 1 1 888450282
1 2`
	expect := 14
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 0 657276545
1 2`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 5 0 10000
1 2 3 4`
	expect := 1825
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 0 8 78731972
1 52 76 81`
	expect := 1108850
	runSample(t, s, expect)
}
