package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectInf bool, expectRes []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	inf, res := drive(reader)

	if expectInf != inf {
		t.Fatalf("sample expect inf %t, but got %t", expectInf, inf)
	}
	if expectInf {
		return
	}

	if !slices.Equal(expectRes, res) {
		t.Fatalf("sample expect res %v, but got %v", expectRes, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
4 1 7
`
	expectInf := false
	expectRes := []int{-2, 10}
	runSample(t, s, expectInf, expectRes)
}

func TestSample2(t *testing.T) {
	s := `1
10
`
	expectInf := true
	runSample(t, s, expectInf, nil)
}

func TestSample3(t *testing.T) {
	s := `4
1 3 5 9
`
	expectInf := false
	expectRes := []int{7}
	runSample(t, s, expectInf, expectRes)
}

func TestSample4(t *testing.T) {
	s := `4
4 3 4 5
`
	expectInf := false
	runSample(t, s, expectInf, nil)
}

func TestSample5(t *testing.T) {
	s := `2
2 4
`
	expectInf := false
	expectRes := []int{0, 3, 6}
	runSample(t, s, expectInf, expectRes)
}
