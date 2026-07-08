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
	if len(res) != 1 || !slices.Equal(res[0], expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5
5 1 4 2 3
`, []int{10, 6, 3, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3
3 2 1
`, []int{3, 0})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4
3 1 2 4
`, []int{6, 2, 0})
}

func TestDisconnectedComponentsAtLowerThreshold(t *testing.T) {
	runSample(t, `1
5
1 3 2 5 4
`, []int{10, 2, 1, 0})
}
