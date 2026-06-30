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
	runSample(t, `5
1 1 0 3
1 1 1 2
1 1 2 1
1 1 3 0
2 1 3 0
`, []int{1, 2, 3, 5})
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 1 0 3
10 1 2 1
2 2 1 1
10 1 1 2
3 1 3 0
`, []int{1, 3, 5})
}

func TestNoSwappedReconstruction(t *testing.T) {
	runSample(t, `3
1 1 0 2
10 2 1 0
5 1 0 2
`, []int{1, 2})
}
