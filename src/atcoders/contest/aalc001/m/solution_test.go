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
	s := `4 3
5 3 12 10
2 1 4
2 2 3
2 3 3
`
	expect := []int{1073741821, 1073741823, 12}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 6
7 10 12 3 5
2 2 4
1 3 6
2 1 5
2 3 3
1 1 1073741823
2 1 2
`
	expect := []int{1073741820, 1073741819, 6, 1073741813}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 5
1073741823 0 536870912
2 1 2
2 1 3
1 2 1073741823
2 1 3
2 2 2
`
	expect := []int{1073741823, 536870911, 1073741823, 1073741823}
	runSample(t, s, expect)
}
