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
	runSample(t, `5 5
2 0 1 2 0
2 1 5
2 2 4
1 2 2
2 1 3
2 2 5
`, []int{5, 0, 2, 4})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 6
0 2 1 0 2 1
2 1 6
1 1 2
2 1 4
1 4 2
2 3 6
2 4 4
`, []int{5, 5, 2, 0})
}
