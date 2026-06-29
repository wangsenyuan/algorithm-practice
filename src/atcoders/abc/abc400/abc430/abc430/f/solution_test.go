package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
RLLR
`, []int{2, 4, 3, 4, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
RL
`, []int{2, 2, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `2
L
`, []int{1, 1})
}

func TestSample4(t *testing.T) {
	runSample(t, `3
RR
`, []int{1, 1, 1})
}

func TestSample5(t *testing.T) {
	runSample(t, `20
RLLLLLLLLRLRRLLLRLR
`, []int{5, 9, 13, 14, 15, 17, 18, 19, 19, 20, 20, 19, 19, 18, 17, 16, 14, 12, 9, 5})
}
