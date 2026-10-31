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
	runSample(t, `5 8
2 1 3 0 4
2 2 4
3 5 5
1 3 1
2 2 4
3 5 5
2 4 5
3 2 4
2 1 2
`, []int{3, 3, 5, 3, 6, 0, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 6
0
2 1 1
3 1 1
1 1 1000000000
2 1 1000000000
3 1 1000000000
2 1 1000000000000000
`, []int{2, 0, 1, 1, 2})
}
