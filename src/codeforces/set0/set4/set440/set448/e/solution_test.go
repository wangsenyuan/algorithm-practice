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
	s := `10 3`
	expect := []int{1, 1, 1, 2, 1, 1, 5, 1, 1, 2, 1, 5, 1, 2, 5, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2`
	expect := []int{1, 1, 2, 1, 2, 4}
	runSample(t, s, expect)
}
