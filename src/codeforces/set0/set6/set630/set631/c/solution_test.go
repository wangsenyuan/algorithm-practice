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
	s := `3 1
1 2 3
2 2
`
	expect := []int{2, 1, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 2 4 3
2 3
1 2
`
	expect := []int{2, 4, 1, 3}
	runSample(t, s, expect)
}
