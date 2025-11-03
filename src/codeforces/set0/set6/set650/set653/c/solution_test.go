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
	s := `5
2 8 4 7 7
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
200 150 100 50
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
3 2 1 4 1 4 1 4 1 4
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
1 2 3 4 5 6 7 8 9
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
1 1 1 4 3
`
	expect := 1
	runSample(t, s, expect)
}
