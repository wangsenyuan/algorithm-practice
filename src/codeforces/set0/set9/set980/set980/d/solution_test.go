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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
5 5
`
	expect := []int{3, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
5 -4 2 1 8
`
	expect := []int{5, 5, 3, 2, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
0
`
	expect := []int{1}
	runSample(t, s, expect)
}
