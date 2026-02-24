package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, ok, res := drive(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}
	if len(res) > 40 {
		t.Fatalf("Sample result %v, too long", res)
	}

	for _, v := range res {
		for i := range a {
			a[i] = abs(a[i] - v)
		}
	}

	if slices.Max(a) != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
5
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `2
0 0
`
	runSample(t, s, true)
}

func TestSample3(t *testing.T) {
	s := `3
4 6 8
`
	runSample(t, s, true)
}

func TestSample4(t *testing.T) {
	s := `4
80 40 20 10
`
	runSample(t, s, true)
}

func TestSample5_Impossible(t *testing.T) {
	s := `5
1 2 3 4 5
`
	runSample(t, s, false)
}
