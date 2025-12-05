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
func TestSamle1(t *testing.T) {
	s := `4 2
1 1 2 2
1 3 2
1 4 2
`
	expect := []int{1, -1}
	runSample(t, s, expect)
}

func TestSamle2(t *testing.T) {
	s := `5 3
1 2 1 3 2
2 5 3
1 2 3
5 5 2
`
	expect := []int{2, 1, 2}
	runSample(t, s, expect)
}
