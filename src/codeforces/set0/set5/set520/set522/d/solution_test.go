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
	s := `5 3
1 1 2 3 2
1 5
2 4
3 5
`
	expect := []int{1, -1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 5
1 2 1 3 2 3
4 6
1 3
2 5
2 4
1 6
`
	expect := []int{2, 2, 3, -1, 2}
	runSample(t, s, expect)
}
