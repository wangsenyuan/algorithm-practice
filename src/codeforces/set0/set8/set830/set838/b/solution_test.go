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
	s := `5 9
1 3 1
3 2 2
1 4 3
3 5 4
5 1 5
3 1 6
2 1 7
4 1 8
2 1 1
2 1 3
2 3 5
2 5 2
1 1 100
2 1 3
1 8 30
2 4 2
2 2 4
`
	expect := []int{0, 1, 4, 8, 100, 132, 10}
	runSample(t, s, expect)
}
