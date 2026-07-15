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
4 3
1 2 4 7
`, []int{4})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
6 4
1 5 3 6 2 4
`, []int{6})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3 2
2 2 3
`, []int{1})
}
