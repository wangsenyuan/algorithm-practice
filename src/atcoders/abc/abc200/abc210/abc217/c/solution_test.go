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
	runSample(t, `3
2 3 1
`, []int{3, 1, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2 3
`, []int{1, 2, 3})
}

func TestSample3(t *testing.T) {
	runSample(t, `5
5 3 2 4 1
`, []int{5, 3, 2, 4, 1})
}
