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
	s := `5 4
1 2
2 3
2 4
4 5
2 1
2 5
1 2
2 5
`
	expect := []int{0, 3, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 19
10 7
10 6
1 4
7 8
5 1
6 2
7 5
6 3
7 9
2 9
2 8
2 2
2 7
2 6
2 8
2 1
2 6
2 4
2 6
2 2
2 1
2 6
2 1
2 7
2 10
1 10
2 9
2 5
`
	expect := []int{3, 3, 5, 2, 4, 3, 0, 4, 1, 4, 5, 0, 4, 0, 2, 3, 2, 1}
	runSample(t, s, expect)
}
