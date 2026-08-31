package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `12 4
0 0 1 1 0 1 0 1 0 1 1 0
1 12
2 7
5 10
6 11
`, []int{4, 2, 3, -1})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 3
0 0 0 1 1 1
1 3
4 6
1 6
`, []int{1, 1, 2})
}
