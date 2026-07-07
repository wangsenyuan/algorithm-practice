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
	runSample(t, `3 4 5
1 2
1 3
3 4
3 1
2 2
5
1 1
1 2
2 2
2 4
1 2
`, []int{2, 1, 0, 1, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, `1 2 1
1 2
7
2 1
2 1
2 1
2 1
2 1
2 1
2 1
`, []int{0, 0, 0, 0, 0, 0, 0})
}

func TestSample3(t *testing.T) {
	runSample(t, `4 4 16
1 1
1 2
1 3
1 4
2 1
2 2
2 3
2 4
3 1
3 2
3 3
3 4
4 1
4 2
4 3
4 4
7
2 1
1 1
2 2
1 2
2 3
1 3
2 4
`, []int{4, 3, 3, 2, 2, 1, 1})
}
