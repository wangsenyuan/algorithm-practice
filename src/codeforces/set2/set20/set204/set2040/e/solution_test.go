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
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
4 4
1 2
2 3
2 4
2 0
3 0
4 0
3 1
`, []int{1, 6, 6, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
12 10
1 2
2 3
2 4
1 5
5 6
6 7
6 8
6 9
8 10
10 11
10 12
6 0
9 0
10 0
11 0
3 1
7 1
10 1
12 1
12 2
11 12
`, []int{4, 9, 8, 15, 2, 3, 6, 9, 5, 5})
}
