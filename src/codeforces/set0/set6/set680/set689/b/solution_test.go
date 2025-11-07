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
	n := len(res)
	expect := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &expect[i])
	}
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 2 3
0 1 2
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
0 1 2 3 4
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `7
4 4 4 4 7 7 7
0 1 2 1 2 3 3 
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `98
17 17 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 57 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 87 90 90 90 90 90 90 90 90 90 90 90 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 92 95 95 95 95 95 97 98 98
0 1 2 3 4 5 6 7 8 8 7 6 5 4 3 2 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 21 20 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 4 4 5 6 5 6 7 8 
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `91
4 6 23 23 23 23 23 28 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 39 47 47 47 54 54 54 54 54 54 54 58 58 58 58 58 58 69 69 69 69 69 69 69 69 69 69 69 69 70 70 70 70 70 70 70 70 70 70 71 72 72 72 73 75 77 77 77 82 82 84 84 84 84 84 85 86 87 89 89 90 91
0 1 2 1 2 2 3 4 5 6 7 8 9 10 10 9 8 7 6 5 4 3 2 3 4 5 6 5 6 7 8 9 9 8 7 6 5 4 3 4 5 6 7 8 9 10 9 10 9 8 7 6 5 4 5 6 7 6 7 8 9 10 11 10 9 8 7 6 5 6 6 7 8 9 10 11 11 12 13 14 14 13 14 14 15 16 17 18 19 20 21 
`
	runSample(t, s)
}
