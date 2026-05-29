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
	s := `3 3
1 1 2 4
1 2 3 5
2 1 3
`
	expect := []int{8}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
1 1 3 4
2 1 1
2 2 2
2 3 3
`
	expect := []int{3, 2, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 6
1 1 5 3
1 2 7 9
1 10 10 11
1 3 8 12
1 1 10 3
2 1 10
`
	expect := []int{129}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 10
1 5 9 60144710
2 3 3
2 3 4
2 6 10
1 8 9 11764737
1 1 9 38454635
1 3 10 58013969
1 2 3 49601991
1 1 3 73003418
2 4 6
`
	expect := []int{0, 0, 240578810, 260802192}
	runSample(t, s, expect)
}
