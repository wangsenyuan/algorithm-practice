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
	s := `4 3 4
H 2
V 2
V 3
V 1
`
	expect := []int{8, 4, 4, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 6 5
H 4
V 3
V 5
H 2
V 1
`
	expect := []int{28, 16, 12, 6, 4}
	runSample(t, s, expect)
}
