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
	s := `3 0`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
2 > 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
3 = 6
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 5
17 <= 10
16 >= 18
9 > 18
8 = 8
6 >= 13
`
	expect := 6804
	runSample(t, s, expect)
}