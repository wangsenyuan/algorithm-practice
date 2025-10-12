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
	s := `7 2
1 2
1 3
1 4
3 5
3 6
3 7
2 7
`
	expect := []int{2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 4
1 2
2 3
2 4
4 5
4 6
2 4 5 6
`
	expect := []int{2, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 2
1 16
12 5
15 19
18 9
8 4
10 16
9 16
20 15
14 19
7 4
18 12
17 12
2 20
6 14
3 19
7 19
18 15
19 13
9 11
12 18
`
	expect := []int{12, 1}
	runSample(t, s, expect)
}
