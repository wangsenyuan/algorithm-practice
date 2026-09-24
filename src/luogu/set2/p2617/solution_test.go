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
	runSample(t, `5 3
3 2 1 4 7
Q 1 4 3
C 2 6
Q 2 5 3
`, []int{3, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1
0
Q 1 1 1
`, []int{0})
}
