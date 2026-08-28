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
	runSample(t, `6 4
2 3
4 5
6 6
1 1
`, []int{6, 11, 14, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `5 2
2 2
3 4
`, []int{3, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, `8 2
3 1
5 6
`, []int{3, 15})
}
