package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
0 0
1 1
2 3
-5 3
`
	expect := []int{14, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
5 5
1000 1000
-1000 1000
3 100
`
	expect := []int{1995, 1995}
	runSample(t, s, expect)
}
