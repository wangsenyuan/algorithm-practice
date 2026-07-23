package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
4
1 2 3 4
5
5 3 1 5 2
6
3 4 2 6 1 5
7
1 2 1 4 2 3 5
`, [][]int{
		{6, 6, 7, 9},
		{17, 16, 14, 14, 17},
		{21, 21, 20, 20, 21, 21},
		{17, 17, 17, 17, 21, 21, 22},
	})
}
