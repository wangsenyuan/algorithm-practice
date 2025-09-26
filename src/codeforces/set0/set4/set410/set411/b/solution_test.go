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
	s := `4 3 5
1 0 0
1 0 2
2 3 1
3 2 0
`
	expect := []int{1, 1, 3, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2 2
1 2
1 2
2 2
`
	expect := []int{1, 1, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1 2
1
2
`
	expect := []int{0, 0}
	runSample(t, s, expect)
}
