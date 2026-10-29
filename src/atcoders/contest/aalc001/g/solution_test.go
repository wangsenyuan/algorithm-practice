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
	runSample(t, `4 5
1 2 3 4
2 1 4 1
2 2 4 2
1 2 3 5
2 1 3 3
2 2 4 4
`, []int{10, 29, 856, 6753})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 5
998244352 0 2
2 1 3 2
1 1 2 1
2 1 3 4
1 2 3 1000000000
2 1 3 1
`, []int{5, 17, 3511297})
}
