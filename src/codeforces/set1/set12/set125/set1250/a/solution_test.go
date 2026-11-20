package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5
3 2 1 3 3
`, [][]int{{1, 2}, {2, 3}, {1, 3}})
}

func TestSample2(t *testing.T) {
	runSample(t, `10 6
7 3 5 7 3 6
`, [][]int{{1, 2}, {2, 3}, {1, 3}, {4, 7}, {4, 5}, {6, 7}, {5, 7}, {8, 8}, {9, 9}, {10, 10}})
}
