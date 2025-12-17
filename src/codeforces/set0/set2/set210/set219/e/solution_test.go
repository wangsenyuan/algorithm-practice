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
	s := `7 11
1 15
1 123123
1 3
1 5
2 123123
2 15
1 21
2 3
1 6
1 7
1 8
	`
	runSample(t, s, []int{1, 7, 4, 2, 7, 4, 1, 3})
}
