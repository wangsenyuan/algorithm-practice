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
	runSample(t, `13 23
1 2
1 3
2 4
2 5
3 6
3 7
4 8
4 9
6 10
6 11
7 12
7 13
2 1
0 8 13 3
0 9 12 5
2 9
2 8
2 2
0 10 12 1
2 2
1 3
2 7
2 1
0 9 5 6
2 4
2 5
1 7
0 9 12 4
0 10 5 7
2 1
2 4
2 12
1 2
2 5
2 3
`, []int{-1, 3, 5, -1, 1, -1, 1, 1, 3, 6, 7, 7, 4, 6})
}
