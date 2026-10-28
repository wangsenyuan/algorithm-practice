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
	runSample(t, `8
2 1 10
1 3 7
2 1 10
1 5 12
2 4 9
2 3 12
1 1 4
2 1 12
`, []int{0, 5, 3, 7, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, `6
1 999999999999999998 1000000000000000000
2 999999999999999997 1000000000000000000
1 1 1000000000000000000
2 1 1000000000000000000
2 999999999999999997 1000000000000000000
2 1 1
`, []int{3, 999999999999999997, 1, 1})
}
