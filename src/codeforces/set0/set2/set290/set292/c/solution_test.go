package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	slices.Sort(res)
	slices.Sort(expect)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
0 1 2 9 8 7
`
	expect := []string{
		"78.190.209.187",
		"79.180.208.197",
		"87.190.209.178",
		"89.170.207.198",
		"97.180.208.179",
		"98.170.207.189",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
4
`
	expect := []string{
		"4.4.4.4",
		"4.4.4.44",
		"4.4.44.4",
		"4.4.44.44",
		"4.44.4.4",
		"4.44.4.44",
		"4.44.44.4",
		"4.44.44.44",
		"44.4.4.4",
		"44.4.4.44",
		"44.4.44.4",
		"44.4.44.44",
		"44.44.4.4",
		"44.44.4.44",
		"44.44.44.4",
		"44.44.44.44",
	}
	runSample(t, s, expect)
}