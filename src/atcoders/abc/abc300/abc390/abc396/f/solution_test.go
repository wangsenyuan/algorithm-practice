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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
2 1 0
`, []int{3, 1, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 6
5 3 5 0 1
`, []int{7, 3, 3, 1, 1, 5})
}

func TestSample3(t *testing.T) {
	runSample(t, `7 7
0 1 2 3 4 5 6
`, []int{0, 6, 10, 12, 12, 10, 6})
}
