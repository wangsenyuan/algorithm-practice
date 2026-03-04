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
	s := `5
1 2 2 3 3
3
2 5 3
2 5 2
1 5 5
`
	expect := []int{2, 3, 1}
	runSample(t, s, expect)
}
