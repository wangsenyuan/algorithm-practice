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
	s := `5 5
1 2
2 4
2 5
4 3
4 5
`
	expect := []int{1, 2, -1, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
1 1
1 2
3 1
3 1
3 5
`
	expect := []int{1, 0, -1, -1, -1}
	runSample(t, s, expect)
}
