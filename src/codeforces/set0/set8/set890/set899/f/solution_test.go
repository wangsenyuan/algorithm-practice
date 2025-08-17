package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 2
abac
1 3 a
2 2 c
`
	expect := "b"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
A0z
1 3 0
1 1 z
`
	expect := "Az"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 4
agtFrgF4aF
2 5 g
4 9 F
1 5 4
1 7 a
`
	expect := "tFrg4"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9 5
aAAaBBccD
1 4 a
5 6 c
2 3 B
4 4 D
2 3 A
`
	expect := "AB"
	runSample(t, s, expect)
}