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
	s := `7 4
1 3
1 2
2 0 1
2 1 2
`
	expect := []int{4, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 9
2 2 9
1 1
2 0 1
1 8
2 0 8
1 2
2 1 3
1 4
2 2 4
`
	expect := []int{7, 2, 10, 4, 5}
	runSample(t, s, expect)
}
