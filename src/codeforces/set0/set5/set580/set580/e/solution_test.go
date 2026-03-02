package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1 2
112
2 2 3 1
1 1 3 8
2 1 2 1
`
	expect := []bool{false, true}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2 3
334934
2 2 5 2
1 4 4 3
2 1 6 3
1 2 3 8
2 3 6 1
`
	expect := []bool{false, true, false}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 1 2
34075930750342906718
2 1 20 20
1 1 20 6
2 1 20 1
`
	expect := []bool{true, true}
	runSample(t, s, expect)
}
