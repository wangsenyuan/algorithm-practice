package main

import (
	"bufio"
	"sort"
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
	a = append(a, res...)
	if len(a) != 4 {
		t.Fatalf("Sample expect 4 numbers, but got %d", len(a))
	}
	sort.Ints(a)
	if !check(a) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1
1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1
1
1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1
2
2
3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
472
107
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `2
3
13
`
	expect := false
	runSample(t, s, expect)
}
