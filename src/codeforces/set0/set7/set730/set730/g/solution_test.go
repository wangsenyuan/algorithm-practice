package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
9 2
7 3
2 4
`, [][]int{
		{9, 10},
		{1, 3},
		{4, 7},
	})
}

func TestSample2(t *testing.T) {
	runSample(t, `4
1000000000 1000000
1000000000 1000000
100000000 1000000
1000000000 1000000
`, [][]int{
		{1000000000, 1000999999},
		{1, 1000000},
		{100000000, 100999999},
		{1000001, 2000000},
	})
}
