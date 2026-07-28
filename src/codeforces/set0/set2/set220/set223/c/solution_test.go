package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 1
1 2 3
`, []int{1, 3, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 0
3 14 15 92 6
`, []int{3, 14, 15, 92, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, `5 20
11 5 6 8 11
`, []int{11, 225, 2416, 18118, 106536})
}
