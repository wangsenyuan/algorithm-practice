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
	runSample(t, `5 1
1 2
1 3
3 5
3 4
2 3
`, []int{0, 3, 3, 3, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `11 3
1 2
2 3
2 4
1 5
5 6
5 7
5 8
6 9
8 10
8 11
2 9
3 6
2 8
`, []int{0, 6, 7, 6, 0, 2, 0, 5, 4, 5, 5})
}
