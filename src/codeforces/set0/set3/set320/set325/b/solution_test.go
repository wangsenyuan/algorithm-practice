package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	expect := []int{3, 4}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 25
	expect := []int{20}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	runSample(t, n, nil)
}
