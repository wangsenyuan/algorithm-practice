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
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 0 4
3 1 4 2
`, []int{1, 4, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1 2
3 2 1
`, []int{2, 3, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `5 1 3
1 3 5 2 4
`, []int{1, 2, 3, 5, 4})
}

func TestSample4(t *testing.T) {
	runSample(t, `2 0 1
1 2
`, []int{1, 2})
}

func TestSample5(t *testing.T) {
	runSample(t, `6 3 4
1 3 4 6 5 2
`, []int{1, 3, 4, 5, 2, 6})
}
