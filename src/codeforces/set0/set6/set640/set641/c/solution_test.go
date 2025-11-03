package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)

	if !slices.Equal(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
1 2
2
1 2
`

	// 1 2 3 4 5 6
	// 5 6 1 2 3 4
	// 6 5 2 1 4 3

	expect := []int{4, 3, 6, 5, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 1
2
1 -2
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 2
2
1 3
`
	expect := []int{1, 4, 3, 2}
	runSample(t, s, expect)
}
