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
	s := `5
2 1 4 3 5
`
	expect := []int{3, 2, 2, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 3 4
`
	expect := []int{3, 2, 1, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
1 9 6 5 2 7 10 4 8 3
`
	expect := []int{2, 3, 3, 3, 2, 1, 2, 1, 1, 0}
	runSample(t, s, expect)
}
