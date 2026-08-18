package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10
100 108 114 118 120 5 7 19 13 11
`, []int{100, 102, 106, 112, 120, 5, -1, -3, -1, 11})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3
`, []int{1, 2, 3})
}

func TestSample3(t *testing.T) {
	runSample(t, `4
10 10 8 4
`, []int{10, 6, 4, 4})
}
