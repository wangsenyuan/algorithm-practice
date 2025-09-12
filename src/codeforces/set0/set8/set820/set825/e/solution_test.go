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
	s := `3 3
1 2
1 3
3 2
`
	expect := []int{1, 3, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
3 1
4 1
2 3
3 4
2 4
`
	expect := []int{4, 1, 2, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 4
3 1
2 1
2 3
4 5
`
	expect := []int{3, 1, 2, 4, 5}
	runSample(t, s, expect)
}
