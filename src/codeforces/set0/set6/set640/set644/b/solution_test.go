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
	s := `5 1
2 9
4 8
10 9
15 2
19 1
`
	expect := []int{11, 19, -1, 21, 22}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1
2 8
4 8
10 9
15 2
`
	expect := []int{10, 18, 27, -1}
	runSample(t, s, expect)
}
