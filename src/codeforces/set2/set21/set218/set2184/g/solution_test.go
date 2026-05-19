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
	runSample(t, `5 5
1 2 3 4 5
2 1 5
1 1 5
1 2 5
1 3 1
2 1 5`, []int{1, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
5 5 1
2 3 3`, []int{0})
}
