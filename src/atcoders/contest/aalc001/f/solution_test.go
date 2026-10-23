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
	runSample(t, `8 7
1 1 2 2 2 3 2 2
2 1 8
2 4 8
1 6 2
2 3 8
1 4 7
2 1 8
2 4 4
`, []int{3, 2, 6, 4, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `6 7
0 1000000000 1000000000 0 0 0
2 1 6
2 2 5
1 3 0
2 2 6
1 2 0
2 1 6
1 6 1000000000
`, []int{3, 2, 4, 6})
}
