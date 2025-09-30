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
	s := `4
1 2 3 4
1 2
2 3
2 4
`
	expect := []int{10, 9, 3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `15
1 2 3 1 2 3 3 1 1 3 2 2 1 2 3
1 2
1 3
1 4
1 14
1 15
2 5
2 6
2 7
3 8
3 9
3 10
4 11
4 12
4 13
`
	expect := []int{6, 5, 4, 3, 2, 3, 3, 1, 1, 3, 2, 2, 1, 2, 3}
	runSample(t, s, expect)
}
