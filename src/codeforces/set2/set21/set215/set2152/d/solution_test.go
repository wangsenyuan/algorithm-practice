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
4 3 2 5 6
1 1
1 2
2 4
3 5
1 5
`
	expect := []int{2, 3, 5, 6, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 1
314 159 265 358 979 323 846 264 338 327
1 10
`
	expect := []int{91}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1
9 9 9
1 3
`
	expect := []int{10}
	runSample(t, s, expect)
}
