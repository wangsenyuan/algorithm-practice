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
	if res == nil {
		res = []int{}
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10
`, []int{5, 10})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
`, []int{})
}

func TestSample3(t *testing.T) {
	runSample(t, `50
`, []int{5, 10, 13, 17, 20, 25, 26, 29, 34, 37, 40, 41, 45, 50})
}
