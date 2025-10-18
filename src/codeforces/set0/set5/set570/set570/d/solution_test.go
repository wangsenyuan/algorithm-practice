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
	s := `6 5
1 1 1 3 3
zacccd
1 1
3 3
4 1
6 1
1 2
`
	expect := []bool{true, false, true, true, true}
	runSample(t, s, expect)
}
