package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x string, y string, expect []int) {
	res := solve(x, y)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "AbC", "DCbA", []int{3, 0})
}
