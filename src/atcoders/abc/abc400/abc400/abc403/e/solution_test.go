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
	s := `4
1 at
2 watcoder
2 atcoder
1 wa
`
	expect := []int{0, 1, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
1 w
1 avko
2 atcoder
1 bzginn
2 beginner
1 atco
2 contest
1 ntxcdg
1 atc
1 contest
`

// (w, avko, gzginn, atco, ntxcdg, atc, contest)
// (*atcoder, beginner, *contest)

	expect := []int{0, 0, 1, 1, 2, 1, 2, 2, 2, 1}
	runSample(t, s, expect)
}
