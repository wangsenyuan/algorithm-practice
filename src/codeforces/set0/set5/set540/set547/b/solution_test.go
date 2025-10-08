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
	s := `10
1 2 3 4 5 4 3 2 1 6
`
	expect := []int{6, 4, 4, 3, 3, 2, 2, 1, 1, 1}
	runSample(t, s, expect)
}
