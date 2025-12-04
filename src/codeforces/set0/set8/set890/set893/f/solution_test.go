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
	s := `5 2
1 3 2 3 5
2 3
5 1
3 4
4 1
2
1 2
2 3`

	runSample(t, s, []int{2, 5})
}

func TestSample2(t *testing.T) {
	s := `3 3
1 3 2
1 2
1 3
2
2 3
1 1`

	runSample(t, s, []int{2, 1})
}
