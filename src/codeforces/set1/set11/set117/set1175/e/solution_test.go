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
	s := `2 3
1 3
2 4
1 3
1 4
3 4`
	runSample(t, s, []int{1, 2, 1})
}

func TestSample2(t *testing.T) {
	s := `3 4
1 3
1 3
4 5
1 2
1 3
1 4
1 5
`
	runSample(t, s, []int{1, 1, -1, -1})
}
