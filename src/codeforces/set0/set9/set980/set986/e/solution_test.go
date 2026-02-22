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
	s := `4
1 2
1 3
1 4
6 4 9 5
3
2 3 6
2 3 2
3 4 7`
	expect := []int{36, 4, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2
2 3
2 4
1 5
5 6
100000 200000 500000 40000 800000 250000
3
3 5 10000000
6 2 3500000
4 1 64000
`
	expect := []int{196000, 12250, 999998215}
	runSample(t, s, expect)
}
