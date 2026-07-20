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
	runSample(t, `1
3 3
1 -2 3
4 -5 2
1 6 -1
`, []int{3})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
2 4
-1 -1 -1 1
-1 -1 -1 -1
`, []int{-5})
}
