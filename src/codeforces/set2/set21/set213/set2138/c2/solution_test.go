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
	s := `7 3
1 1 2 2 3 3
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 2
1 1 2 3 1 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 0
1 2 3 4
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 2
1 1 1 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 4
1 1 1 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestExample2_1(t *testing.T) {
	s := `2 0
1
`
	expect := 2
	runSample(t, s, expect)
}

func TestExample2_2(t *testing.T) {
	s := `2 1
1
`
	expect := 2
	runSample(t, s, expect)
}

func TestExample2_3(t *testing.T) {
	s := `3 0
1 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestExample2_4(t *testing.T) {
	s := `3 1
1 2
`
	expect := 3
	runSample(t, s, expect)
}

func TestExample2_5(t *testing.T) {
	s := `3 1
1 1
`
	expect := 2
	runSample(t, s, expect)
}

func TestExample7(t *testing.T) {
	s := `12 7
1 1 1 2 3 4 5 6 7 7 7
`
	expect := 4
	runSample(t, s, expect)
}
