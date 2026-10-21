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
	_, res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
101
`, []int{2})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
100
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `4
0000
`, nil)
}

func TestSample4(t *testing.T) {
	runSample(t, `4
1010
`, []int{1, 3})
}

func TestSample5(t *testing.T) {
	runSample(t, `2
01
`, []int{1})
}
