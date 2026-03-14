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
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3 2
Sba
ccc
aac
ccc
abT
`
	expect := "bcccc"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 1
Sxyy
yxxx
yyyT
`
	expect := "xxxx"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 3 3
TyS
`
	expect := "y"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 4 1
SxyT
`
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 4 1
SbbT
aaaa
abba
`
	expect := "bb"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 5 2
SbcaT
acbab
acccb
`
	expect := "aacccaa"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `3 3 1
aaa
aaa
TSa
`
	expect := ""
	runSample(t, s, expect)
}

func TestNextPermuatation(t *testing.T) {
	tests := []struct {
		state  int
		expect int
	}{
		{0, 0},
		{1, 2},
		{3, 5},
		{5, 6},
		{6, 9},
		{7, 11},
		{10, 12},
	}

	for _, tc := range tests {
		got := nextPermuatation(tc.state)
		if got != tc.expect {
			t.Fatalf("nextPermuatation(%d) expect %d, but got %d", tc.state, tc.expect, got)
		}
	}
}
