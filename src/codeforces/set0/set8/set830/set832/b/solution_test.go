package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ab
a?a
2
aaa
aab
`
	expect := []bool{true, false}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abc
a?a?a*
4
abacaba
abaca
apapa
aaaaax
`
	expect := []bool{false, true, false, true}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `y
b*b
1
b
`
	expect := []bool{false}
	runSample(t, s, expect)
}
