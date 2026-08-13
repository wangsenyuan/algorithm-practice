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
	runSample(t, `3
5 1 2 3 4 6
3 2 5 1
4 1 9 2 3
`, []int{1, 5, 2, 3, 9, 6, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `2
2 1 6
1 6
`, []int{6, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
3 6 1 1
`, []int{1, 6})
}

func TestSample4(t *testing.T) {
	runSample(t, `5
4 2 3 3 4
5 1 2 4 3 1
2 4 1
3 3 3 1
5 4 3 2 2 2
`, []int{1, 3, 2, 4})
}

func TestSample5(t *testing.T) {
	runSample(t, `5
4 2 3 1 4
5 2 5 5 6 5
5 3 4 7 5 5
8 3 6 4 3 1 1 5 4
2 1 1
`, []int{1, 4, 3, 2, 5, 6, 7})
}
