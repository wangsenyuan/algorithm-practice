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
	runSample(t, `8 3
1 5
2 7
3 4
`, []string{"Yes", "No", "Yes"})
}

func TestSample2(t *testing.T) {
	runSample(t, `999999 4
123456 987654
888888 999999
1 3
2 777777
`, []string{"Yes", "No", "Yes", "No"})
}
