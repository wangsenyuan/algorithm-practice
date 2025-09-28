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
	s := `3
0 3
1 3
3 8
`
	expect := []int{6, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 3
2 4
5 7
`
	expect := []int{5, 2, 0}
	runSample(t, s, expect)
}
