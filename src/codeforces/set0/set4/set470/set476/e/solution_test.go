package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, p string, expect []int) {
	res := solve(s, p)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aaaaa", "aa", []int{2, 2, 1, 1, 0, 0})
}

func TestSample2(t *testing.T) {
	runSample(t, "axbaxxb", "ab", []int{0, 1, 1, 2, 1, 1, 0, 0})
}

func TestSample3(t *testing.T) {
	runSample(t, "ababababababababa", "aba", []int{4, 4, 4, 4, 4, 4, 3, 3, 3, 2, 2, 2, 1, 1, 1, 0, 0, 0})
}
