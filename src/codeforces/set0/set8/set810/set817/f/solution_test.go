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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3 4
3 1 6
2 1 3
`
	expect := []int{1, 3, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 1 3
3 5 6
2 4 4
3 1 6
`
	expect := []int{4, 4, 4, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
2 9 59
2 8 72
3 19 41
1 1 62
1 50 74
1 53 66
2 59 69
3 66 77
2 62 63
3 57 69
`
	expect := []int{1,
		1,
		1,
		63,
		75,
		75,
		59,
		59,
		59,
		57}
	runSample(t, s, expect)
}
