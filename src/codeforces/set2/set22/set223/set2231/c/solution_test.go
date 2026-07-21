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
3 2 4
`, []int{3})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
7
3 6 7 16 8 8 7
`, []int{11})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3
1 4 2
`, []int{2})
}

func TestSample4(t *testing.T) {
	runSample(t, `1
5
10 10 10 10 10
`, []int{0})
}

func TestSample5(t *testing.T) {
	runSample(t, `1
6
1 1 3 1 1 1
`, []int{3})
}
