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
	s := `4 3 100 30
12 2 6 1
`
	expect := 220
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3 30 100
12 2 6 1
`
	expect := 190
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2 987 789
-8 42 -4 -65 -8 -8
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 6 1 1000
1 1 1 1 1 2
`
	expect := 4005
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 7 3 10
-35 -33 10 15 20 25 30 31
`
	expect := 127
	runSample(t, s, expect)
}
