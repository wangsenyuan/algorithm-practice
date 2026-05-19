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
	runSample(t, `2 2
1 3 5 7
1 1
4 8
`, []int{1, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 2
1 3
1 4
1 2
`, []int{0, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `3 4
1 8 3 10 2 5 7 1
5 3
8 12
1 9
2 1
`, []int{5, 0, 2, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, `2 1
1 2 3 4
3 1
`, []int{1})
}
