package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	res = append(res, n)
	slices.Sort(res)

	if res[0]*res[0]+res[1]*res[1] != res[2]*res[2] {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "3"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "67"
	expect := true
	runSample(t, s, expect)
}