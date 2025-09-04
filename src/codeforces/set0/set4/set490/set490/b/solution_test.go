package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
92 31
0 7
31 0
7 141
`, []int{92, 7, 31, 141})
}

func TestSample2(t *testing.T) {
	runSample(t, `3
0 2
1 3
2 0
`, []int{1, 2, 3})
}
