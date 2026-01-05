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
	s := `2 3 3
1 1 1
3 2
4 0
`
	expect := []int{1, 4, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2 6
3 2
2 2 2
3 3
3 2
2 2 2
3 2
`
	expect := []int{2, 1, 3, 3, 2, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2 2
3 2
2 2 1
`
	expect := []int{2, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `19 16 14
3 10
3 8
3 3
2 5 3
1 7 15
4 0
2 11 2
1 16 3
1 16 3
3 13
1 13 3
4 9
1 5 11
3 1
`
	expect := []int{16, 32, 48, 48, 49, 0, 0, 1, 1, 17, 17, 1, 2, 18}
	runSample(t, s, expect)
}
