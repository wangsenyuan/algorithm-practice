package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if !slices.Equal(expect, ans) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `1 3
0
0
2
4`
	expect := []int{0, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
1 3
0
3`
	expect := []int{2, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
1000000000 1000000000
1000000000`
	expect := []int{31}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1
2
0`
	expect := []int{1}
	runSample(t, s, expect)
}