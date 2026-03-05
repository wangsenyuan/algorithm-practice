package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `abc
bc
cbc
`
	expect := []int{3, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abacaba
abac
abcd
`
	expect := []int{11, 2, 0, 0}
	runSample(t, s, expect)
}
