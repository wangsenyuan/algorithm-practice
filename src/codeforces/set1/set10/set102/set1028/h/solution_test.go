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
	s := `10 10
34 37 28 16 44 36 43 50 22 13
1 3
4 8
6 10
9 10
3 10
8 9
5 6
1 4
1 7
2 6
`
	expect := []int{2,
		0,
		1,
		3,
		0,
		1,
		1,
		1,
		0,
		0,
	}
	runSample(t, s, expect)
}
