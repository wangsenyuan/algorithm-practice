package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `zyxxxxxxyyz
5
5 5
1 3
1 11
1 4
3 6
`
	expect := []bool{true, true, false, true, false}
	runSample(t, s, expect)
}
