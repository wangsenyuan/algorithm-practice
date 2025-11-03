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
	s := `6
1 1 5
3 5 5
1 2 5
3 6 5
2 3 5
3 7 5
	`
	expect := []int{1, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 1 1
2 2 1
3 3 1
	`
	expect := []int{0}
	runSample(t, s, expect)
}
