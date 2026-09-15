package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 3
3 3 3
2 3 3
2 3 1
3 1 1
3 1 2
1 3 1
1 1 2
1 2 2
1 3 2
1 2 1
`, []int{3, 1, 3, 0, 1, 0, 1, 0, 0, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `10 3
2 3 3
3 2 2
3 3 3
3 2 1
2 2 1
1 2 2
3 2 1
2 2 3
2 2 3
3 2 1
`, []int{2, 0, 0, 5, 1, 1, 0, 0, 0, 1})
}
