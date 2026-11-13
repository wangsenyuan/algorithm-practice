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
	runSample(t, `5 2
2 3
4 1
`, []int{2, 3, 2, 2, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 3
2 6
2 5
2 4
`, []int{12, 15, 12, 9, 6, 3})
}
