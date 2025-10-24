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
		t.Errorf("got %v, expect %v", res, expect)
	}
}

func TestSample1(t *testing.T) {
	s := `6
1 5
3 3
4 4
9 2
10 1
12 1
4
1 2
2 4
2 5
2 6
`
	expect := []int{0, 1, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
304 54
88203 83
1
1 2
`
// 88203 - (304 + 54) = 87845
	expect := []int{87845}
	runSample(t, s, expect)
}
