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
	s := `a
b
6
2 7 a
1 3 b
3 7 b
1 9 c
1 1000000000000000000 b
1000000000000000000 1000000000000000000 a
`
	expect := []int{3, 2, 3, 0, 618033988749894848, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `a
b
1
1000000000000000000 1000000000000000000 a
`
	// a, b, ab, bab, abbab, bababbab
	expect := []int{1}
	runSample(t, s, expect)
}
