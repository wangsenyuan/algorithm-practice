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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `xudyhduxyz
3
xyz xyz
dyh xyz
dzy xyz
`
	expect := []int{3, 8, -1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abcabd
3
a c
ab abc
ab d
`
	expect := []int{2, 3, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `baabcabaaa
2
abca baa
aa aba
`
	expect := []int{6, 4}
	runSample(t, s, expect)
}
