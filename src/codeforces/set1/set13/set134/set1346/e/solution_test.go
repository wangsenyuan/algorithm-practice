package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 5 1
3 4
2 1
4 1
3 1
3 1
`, []int{2, 0, 3, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 7 4
3 2
3 2
4 2
3 4
4 1
3 2
5 2
`, []int{2, 2, 0, 3, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `7 15 5
5 3
4 2
6 1
2 4
1 6
3 7
5 6
4 2
6 4
2 6
6 3
6 3
7 6
2 6
7 2
`, []int{-1, 1, 1, 1, 2, 1, 0})
}
