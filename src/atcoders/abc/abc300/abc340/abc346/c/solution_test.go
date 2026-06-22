package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
1 6 3 1
`
	runSample(t, s, 11)
}

func TestSample2(t *testing.T) {
	s := `1 3
346
`
	runSample(t, s, 6)
}

func TestSample3(t *testing.T) {
	s := `10 158260522
877914575 24979445 623690081 262703497 24979445 1822804784 1430302156 1161735902 923078537 1189330739
`
	runSample(t, s, 12523196466007058)
}
