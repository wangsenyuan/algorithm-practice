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
	s := `7 5
8 5 3 2 1 5 6
`
	expect := []int{
		0,
		0,
		1,
		2,
		5,
		1,
		5,
	}
	runSample(t, s, expect)
}
