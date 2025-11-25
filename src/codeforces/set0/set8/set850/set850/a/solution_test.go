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
	s := `6
0 0 0 0 0
1 0 0 0 0
0 1 0 0 0
0 0 1 0 0
0 0 0 1 0
0 0 0 0 1
`
	expect := []int{1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
0 0 1 2 0
0 0 9 2 0
0 0 5 9 0
`
	runSample(t, s, nil)
}
