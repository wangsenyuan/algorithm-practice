package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if len(res) != 1 || !slices.Equal(res[0], expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
1
0 0
`, []int{1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
5
2 3
0 0
4 5
0 0
0 0
`, []int{9, 10, 14, 15, 15})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
7
2 3
4 5
0 0
6 7
0 0
0 0
0 0
`, []int{13, 22, 14, 27, 23, 28, 28})
}
