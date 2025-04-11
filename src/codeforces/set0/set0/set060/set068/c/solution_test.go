package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2 1 2 3
`
	expect := []int{1, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 1 2 3
1 3 0 0 0
2 3 3 4 5
`
	expect := []int{-1, -1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 0 2 1
2 3 0 2 1
1 3 0 2 6
1 4 0 0 1
2 4 0 0 0
3 4 2 3 0
`
	expect := []int{2, 15}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
1 2 0 2 1
1 3 1 2 1
2 3 1 2 1
`
	expect := []int{2, 6}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
1 2 0 5 0
1 3 0 5 0
1 4 0 5 0
1 5 0 5 0
1 6 0 5 0
2 3 0 5 0
2 4 0 5 0
2 5 0 5 0
2 6 0 5 0
3 4 0 5 0
3 5 0 5 0
3 6 0 5 0
4 5 0 5 0
4 6 0 5 0
5 6 1 5 0
`
	expect := []int{1, 5}
	runSample(t, s, expect)
}
