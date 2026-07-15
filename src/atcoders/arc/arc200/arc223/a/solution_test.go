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
3 10
1 2
3 1
9 3
`, []int{5})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3 20
1 2
3 1
9 3
`, []int{6})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4 15
1 1
3 2
5 3
10 4
`, []int{7})
}
