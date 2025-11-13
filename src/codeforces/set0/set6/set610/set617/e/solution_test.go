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
	s := `6 2 3
1 2 1 1 0 3
1 6
3 5
`
	expect := []int{7, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3 1
1 1 1 1 1
1 5
2 4
1 3
`
	expect := []int{9, 4, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 1 0
0 0 0 0 0
1 5
`
	// (1 + 5) * 5 / 2 = 10
	expect := []int{15}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50 2 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
17 35
3 35
`
	expect := []int{190, 561}
	runSample(t, s, expect)
}
