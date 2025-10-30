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
	s := `4 7
5 4 3 2
5 6 5 4
`
	expect := []int{6, 4, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
2 4 4 1 3
1 0 1 2 4
`
	expect := []int{4, 4, 4, 3, 2}
	runSample(t, s, expect)
}
