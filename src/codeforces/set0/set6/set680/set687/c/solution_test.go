package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	var sz int
	fmt.Fscan(reader, &sz)

	expect := make([]int, sz)
	for i := 0; i < sz; i++ {
		fmt.Fscan(reader, &expect[i])
	}

	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 18
5 6 1 10 12 2
16
0 1 2 3 5 6 7 8 10 11 12 13 15 16 17 18 
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 50
25 25 50
3
0 25 50 
`
	runSample(t, s)
}
