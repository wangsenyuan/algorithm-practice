package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 3 2 5 4`
	expect := [][]int{
		{1, 3, 5},
		{2, 4},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
4 3 2 1
`
	expect := [][]int{
		{4},
		{3},
		{2},
		{1},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
10 30 50 101
`
	expect := [][]int{
		{10, 30, 50, 101},
	}
	runSample(t, s, expect)
}
