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
	s := `5 6
1 2
1 3
2 4
3 4
3 5
4 5
`
	expect := []int{0, 1, 2, 1}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11 15
1 2
1 3
2 3
3 4
3 5
4 5
5 6
5 7
6 7
7 8
7 9
8 9
9 10
9 11
10 11
`
	expect := []int{0, 0, 0, 0, 1, 5, 10, 10, 5, 1}

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 18
6 7
4 5
1 7
2 7
1 4
2 5
4 6
2 3
5 6
5 7
1 5
2 4
2 6
1 2
1 3
3 4
1 6
3 5
`
	expect := []int{1, 3, 11, 29, 50, 42}

	runSample(t, s, expect)
}
