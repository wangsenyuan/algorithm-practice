package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8 6
3 2 5
1 2 5
3 2 5
2 4 7
2 1 2
3 1 7
`
	expect := []string{"NO", "YES", "YES"}
	runSample(t, s, expect)
}
