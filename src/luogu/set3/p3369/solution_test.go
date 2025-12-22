package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cmds [][]int, expect []int) {
	res := solve(cmds)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cmds := [][]int{
		{1, 106465},
		{4, 1},
		{1, 317721},
		{1, 460929},
		{1, 644985},
		{1, 84185},
		{1, 89851},
		{6, 81968},
		{1, 492737},
		{5, 493598},
	}
	expect := []int{106465, 84185, 492737}
	runSample(t, cmds, expect)
}
