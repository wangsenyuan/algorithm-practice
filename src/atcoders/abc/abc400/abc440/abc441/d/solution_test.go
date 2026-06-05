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
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 8 3 80 100
1 2 20
1 3 70
2 1 30
2 5 10
3 2 10
3 4 30
3 5 20
5 1 70
`
	expect := []int{1, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 5 3 1 100
1 1 1
2 2 100
1 2 1
1 2 1
1 2 100
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}
