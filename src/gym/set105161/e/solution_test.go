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
	runSample(t, `3 2
2 0 2
2 3 0
1 3 0
`, []int{2, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 6
9 5 0 3 6 7
1 4 7
3 3 233
6 6 0
3 4 4
4 5 15
1 1 0
`, []int{1, 0, 7, 0, 0, 9})
}
