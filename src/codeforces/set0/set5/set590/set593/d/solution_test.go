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
	s := `6 6
1 2 1
1 3 7
1 4 4
2 5 5
2 6 2
1 4 6 17
2 3 2
1 4 6 17
1 5 5 20
2 4 1
1 5 1 3`
	expect := []int{2, 4, 20, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4
1 2 7
1 3 3
3 4 2
3 5 5
1 4 2 100
1 5 4 1
2 2 2
1 1 3 4
`
	expect := []int{2, 0, 2}
	runSample(t, s, expect)
}
