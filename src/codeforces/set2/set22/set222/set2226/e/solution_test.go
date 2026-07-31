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
	runSample(t, `4
0 1 2 3
`, []int{1, 2, 3, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `2
6 7
`, []int{1, 2})
}

func TestSample3(t *testing.T) {
	runSample(t, `6
8 1 7 6 4 3
`, []int{1, 2, 3, 4, 5, 5})
}

func TestSample4(t *testing.T) {
	runSample(t, `9
9 9 8 2 4 4 3 5 3
`, []int{1, 2, 3, 4, 5, 5, 5, 6, 6})
}

func TestInitialElementCanBeReleased(t *testing.T) {
	runSample(t, `2
3 0
`, []int{1, 2})
}

func TestSingleZero(t *testing.T) {
	runSample(t, `1
0
`, []int{1})
}
