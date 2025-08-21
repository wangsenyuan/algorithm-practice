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
		t.Errorf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
-5 20 -3 0
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
-5 20 -3 0
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 6
2 -5 1 3 0 0 -4 -3 1 0
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50 3
6 20 17 19 15 17 3 17 5 16 20 18 9 19 18 18 2 -3 11 11 5 15 4 18 16 16 19 11 20 17 2 1 11 14 18 -8 13 17 19 9 9 20 19 20 19 5 12 19 6 9
`
	expect := 4
	runSample(t, s, expect)
}
