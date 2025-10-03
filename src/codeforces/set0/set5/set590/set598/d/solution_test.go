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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 6 3
******
*..*.*
******
*....*
******
2 2
2 5
4 3
`
	expect := []int{6, 4, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4 1
****
*..*
*.**
****
3 2
`
	expect := []int{8}
	runSample(t, s, expect)
}
