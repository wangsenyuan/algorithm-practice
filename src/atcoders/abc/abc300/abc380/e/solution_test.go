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

func TestSample(t *testing.T) {
	runSample(t, `5 6
1 5 4
1 4 2
2 2
1 3 2
1 2 3
2 3
`, []int{3, 4})
}
