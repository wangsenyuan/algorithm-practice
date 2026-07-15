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
3
1 9 5
`, []int{2})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3
2 2 3
`, []int{0})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
5
11 33 22 55 44
`, []int{3})
}

func TestOddResiduePermutation(t *testing.T) {
	res := solve([]int{1, 5, 7, 8, 9})
	if res != 2 {
		t.Fatalf("expect 2, but got %d", res)
	}
}
