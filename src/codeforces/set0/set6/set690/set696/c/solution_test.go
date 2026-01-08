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

func TestSample1(t *testing.T) {
	s := `1
2
`
	expect := []int{1, 2}
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `3
1 1 1
`
	expect := []int{0, 1}
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `1
983155795040951739
`
	expect := []int{145599903, 436799710}
	runSample(t, s, expect)
}