package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
00110
1 5 1
1 5 2
2 4 1
1 2 0
3 4 0
`, []string{"YES", "YES", "YES", "NO", "NO"})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2
1010
1 4 0
2 3 1
`, []string{"YES", "YES"})
}
