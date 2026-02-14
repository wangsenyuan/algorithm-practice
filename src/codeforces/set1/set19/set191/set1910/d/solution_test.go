package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
4 4 1 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
4 4 1 5 5
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
10 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1 2 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
2 1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4
1 1 1 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `4
1 3 1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5
1 1 3 3 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `29
1 1 3 3 4 6 6 6 8 9 10 11 12 13 14 15 16 16 18 18 20 20 22 23 23 24 26 26 27
`
	expect := true
	runSample(t, s, expect)
}
