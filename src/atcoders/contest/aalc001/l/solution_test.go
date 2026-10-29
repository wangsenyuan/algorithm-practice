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
	runSample(t, `5 7
3 -4 5 -2 1
-1 6 -3 4 -5
2 -7 1 8 -6
4 2
1 2 4
4 1
2 3 5
4 3
3 1 3
4 2
`, []int{3, 1, 4, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 7
0 -3 3 -5 7 -2
-4 2 -1 0 5 -8
6 -2 0 4 -3 1
4 1
4 6
2 2 6
1 3 5
4 3
3 6 6
4 6
`, []int{5, 7, 4, 7})
}
