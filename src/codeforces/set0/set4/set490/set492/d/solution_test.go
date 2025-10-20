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

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 3 2
1
2
3
4
`
	expect := []string{
		"Vanya",
		"Vova",
		"Vanya",
		"Both",
	}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 1
1
2
`
	expect := []string{
		"Both",
		"Both",
	}

	runSample(t, s, expect)
}
