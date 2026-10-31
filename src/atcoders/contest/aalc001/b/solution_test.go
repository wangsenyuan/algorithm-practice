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
5 3 7 9 1
2 2 4
1 2 4 5
2 2 4
1 4 5 10
2 3 5
1 1 5 2
2 1 5
`, []int{3, 8, 11, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 5
1000000000
2 1 1
1 1 1 1000000000
2 1 1
1 1 1 1000000000
2 1 1
`, []int{1000000000, 2000000000, 3000000000})
}
