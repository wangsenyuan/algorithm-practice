package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 10 10
0 12 10
1 6 10
0 1 1
`
	expect := []int{2, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 7 10
3 12 10
5 16 8
`
	expect := []int{0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 8 10
0 12 3
1 1 0
0 3 11
1 6 9
`
	expect := []int{4, 9}
	runSample(t, s, expect)
}
