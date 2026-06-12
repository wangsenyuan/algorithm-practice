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
	s := `<a><b><b></b></b></a><a><b></b><b><v/></b></a><b></b>
4
a
a b b
a b
b a
`
	expect := []int{2, 1, 4, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `<b><aa/></b><aa><b/><b/></aa>
5
aa b
b
aa
b aa
a
`
	expect := []int{2, 3, 2, 1, 0}
	runSample(t, s, expect)
}
