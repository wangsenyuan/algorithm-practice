package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, expect []int) {
	res := solve(m)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 48, []int{9, 42})
}

func TestSample2(t *testing.T) {
	runSample(t, 6, []int{6, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, 994, []int{12, 941})
}

func TestSample4(t *testing.T) {
	runSample(t, 123830583943, []int{17, 123830561521})
}
