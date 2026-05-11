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
	s := `5 3
1 2 3 4 5
1 5 4
1 3 4
3 4 4
	`
	expect := []int{0, -1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 4
3 1 5 2 7 6 4
3 4 2
2 3 5
1 5 6
1 7 3
	`
	expect := []int{2, 0, -1, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
2 1
1 2 1
	`
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1
1
1 1 1
	`
	expect := []int{0}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `7 1
3 4 2 5 7 1 6
1 7 1
	`
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `16 1
16 10 12 6 13 9 14 3 8 11 15 2 7 1 5 4
1 16 4
	`
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `16 1
14 1 3 15 4 5 6 16 7 8 9 10 11 12 13 2
1 16 14
	`
	expect := []int{-1}
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `13 1
12 13 10 9 8 4 11 5 7 6 2 1 3
1 13 2
	`
	expect := []int{-1}
	runSample(t, s, expect)
}
