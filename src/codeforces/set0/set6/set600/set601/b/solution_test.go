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
	s := `10 4
1 5 2 9 1 3 4 2 1 7
2 4
3 8
7 10
1 9
`
	expect := []int{17, 82, 23, 210}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 6
5 7 7 4 6 6 2
1 2
2 3
2 6
1 7
4 7
3 5
`
	expect := []int{2, 0, 22, 59, 16, 8}
	runSample(t, s, expect)
}
