package p3928

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, prices []int, roads [][]int, expect []int) {
	res := minCost(n, prices, roads)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	prices := []int{8, 3}
	roads := [][]int{{0, 1, 1, 2}}
	expect := []int{6, 3}
	runSample(t, n, prices, roads, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	prices := []int{9, 4, 6}
	roads := [][]int{{0, 1, 1, 3}, {1, 2, 4, 2}}
	expect := []int{8, 4, 6}
	runSample(t, n, prices, roads, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	prices := []int{10, 11, 1}
	roads := [][]int{{0, 2, 1, 3}, {1, 2, 3, 4}, {0, 1, 5, 2}}
	expect := []int{5, 11, 1}
	runSample(t, n, prices, roads, expect)
}
