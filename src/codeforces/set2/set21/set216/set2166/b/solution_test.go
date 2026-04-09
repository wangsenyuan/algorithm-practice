package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	expect := make([]int, len(res))
	for i := range len(res) {
		fmt.Fscan(reader, &expect[i])
	}

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
0 2 1
1 2 0
3 6 8
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
0 1 2 3 4
4 3 2 1 0
17 18 20 24 32
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10
5 8 9 3 4 0 2 7 1 6
9 5 1 4 0 3 2 8 7 6
544 768 1024 544 528 528 516 640 516 768
`
	runSample(t, s)
}
