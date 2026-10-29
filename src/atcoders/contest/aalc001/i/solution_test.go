package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 7
4 -5 3 -1 2 -6 5
2 1 7
2 2 6
1 4 10
2 1 7
1 6 1
2 1 7
2 2 2
`, []int64{9, 5, 20, 25, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 5
-8 -3 -5 -2
2 1 4
1 2 7
2 1 4
1 4 6
2 1 4
`, []int64{0, 7, 13})
}

func TestSample3(t *testing.T) {
	runSample(t, `30 1
8 9 3 3 7 -4 5 -9 -2 6 3 -7 2 2 3 -3 -4 9 7 2 -8 9 9 -4 -6 2 3 -9 -10 -10
2 7 28
`, []int64{38})
}

func TestSample4(t *testing.T) {
	runSample(t, `11 1
10 -1 1 -1 10 -100 10 -1 1 -1 10
2 1 9
`, []int64{29})
}

func TestSample5(t *testing.T) {
	runSample(t, `30 1
8 10 6 7 -8 -7 -3 -1 9 -4 -3 -4 -3 -4 -4 7 -2 3 1 10 -5 -2 0 3 -1 8 -6 -4 6 4
2 11 26
`, []int64{29})
}
