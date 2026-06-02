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
		t.Errorf("Sample failed: expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 7
1 2
1 3
2 3
1 4
1 5
2 5
6 7
5
1 2 3 1 5`
	expect := []int{4, 3, 3, 3, 2}
	runSample(t, s, expect)
}
