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
3
12 20 30
`, []int{4})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
6
1 9 1 9 8 1
`, []int{3})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3
6 14 15
`, []int{3})
}
