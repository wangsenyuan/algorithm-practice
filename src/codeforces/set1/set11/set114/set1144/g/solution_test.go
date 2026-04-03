package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	arr, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect {
		return
	}
	var a []int
	var b []int
	for i, v := range arr {
		if res[i] == 0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}
	if !slices.IsSorted(a) {
		t.Fatalf("Sample got wrong a %v", a)
	}
	slices.Reverse(b)
	if !slices.IsSorted(b) {
		t.Fatalf("Sample got wrong b %v", b)
	}
}

func TestSample1(t *testing.T) {
	s := `9
5 1 3 6 8 2 9 0 10
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 4 0 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
2 1 2
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3
1 1 1
`
	expect := false
	runSample(t, s, expect)
}
