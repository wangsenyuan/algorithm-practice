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
	s := `5 5
1 1 1 1 1
1 0 0 0 1
1 0 0 0 1
1 0 0 0 1
1 1 1 1 1
`
	expect := []int{2, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 7
0 0 1 1 1 0 0
0 0 1 0 1 0 0
1 1 1 1 1 1 1
1 0 1 0 1 0 1
1 1 1 1 1 1 1
0 0 1 0 1 0 0
0 0 1 1 1 0 0
`
	expect := []int{2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 1 1
1 1 1
1 1 1
`
	runSample(t, s, nil)
}

func TestSample4(t *testing.T) {
	s := `5 5
0 0 1 0 0
0 0 1 0 0
1 1 1 1 1
0 0 1 0 0
0 0 1 0 0
`
	runSample(t, s, nil)
}

func TestSample5(t *testing.T) {
	s := `5 4
0 0 0 0
0 1 0 0
0 0 0 0
0 0 0 0
0 0 0 0
`
	runSample(t, s, nil)
}

func TestSample6(t *testing.T) {
	s := `5 5
1 1 1 0 0
1 0 1 1 0
1 1 1 1 1
0 0 1 0 1
0 0 1 1 1
`
	runSample(t, s, nil)
}
