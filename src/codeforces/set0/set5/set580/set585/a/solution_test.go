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
	s := `5
4 2 2
4 1 2
5 2 4
3 3 5
5 1 2
`
	expect := []int{1, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
4 5 1
5 3 9
4 1 2
2 1 8
4 1 9
`
	expect := []int{1, 2, 4, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
10 7 10
3 6 11
8 4 10
10 1 11
7 3 13
7 2 13
7 6 14
3 4 17
9 4 20
5 2 24
`
	expect := []int{1, 2, 5}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
5 6 3
7 4 10
9 1 17
2 8 23
9 10 24
6 8 18
3 2 35
7 6 6
1 3 12
9 9 5
`
	expect := []int{1, 2, 3, 4, 5, 7}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10
5 6 1
9 2 6
4 1 5
4 10 5
1 8 23
9 4 21
3 9 6
7 8 34
7 4 24
8 9 21
`
	expect := []int{1, 2, 5, 6, 8}
	runSample(t, s, expect)
}
