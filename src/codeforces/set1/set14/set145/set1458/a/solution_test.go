package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect []int) {
	res := solve(A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 25, 121, 169}
	B := []int{1, 2, 7, 23}
	expect := []int{2, 3, 8, 24}
	runSample(t, A, B, expect)
}
