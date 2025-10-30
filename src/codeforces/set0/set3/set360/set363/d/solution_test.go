package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	r, sum := drive(reader)
	if r != expect[0] || sum != expect[1] {
		t.Errorf("Sample expect %v, but got %v", expect, []int{r, sum})
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 10
5 5
7 6
`
	expect := []int{2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5 2
8 1 1 2
6 3 7 5 2
`
	runSample(t, s, []int{3, 8})
}

func TestSample3(t *testing.T) {
	s := `1 1 2
1
2
`
	runSample(t, s, []int{1, 0})
}
