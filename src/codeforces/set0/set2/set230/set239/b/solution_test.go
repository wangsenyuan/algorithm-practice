package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, cur := range res {
		expect := readNums(reader)
		if !slices.Equal(cur, expect) {
			t.Errorf("Sample expect %v, but got %v", expect, cur)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 4
1>3>22<
1 3
4 7
7 7
1 7
0 1 0 1 0 0 0 0 0 0 
2 2 2 0 0 0 0 0 0 0 
0 0 0 0 0 0 0 0 0 0 
2 3 2 1 0 0 0 0 0 0 
`
	runSample(t, s)
}
