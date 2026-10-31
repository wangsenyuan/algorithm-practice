package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
5 3 7 9 1
2 2 4
1 3 2
2 2 4
1 5 8
2 4 5
`, []int{3, 2, 8})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
0 1000000000 7
2 1 1
2 2 2
1 1 1000000000
2 1 3
`, []int{0, 1000000000, 7})
}
