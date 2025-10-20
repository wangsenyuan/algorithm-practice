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
	s := `6 20
10 50 100 500 1000 5000
8
4200
100000
95000
96000
99000
10100
2015
9950
`
	expect := []int{6,
		20,
		19,
		20,
		-1,
		3,
		-1,
		-1,
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
1 2 3 5 8
8
1
3
5
7
9
11
13
15
`
	expect := []int{1,
		1,
		1,
		2,
		2,
		2,
		2,
		-1,
	}
	runSample(t, s, expect)
}
