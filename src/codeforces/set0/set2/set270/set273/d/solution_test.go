package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 一个4个格子, pow(2, 4) = 16
	// 减去 全0， 对角线 2, 应该就是13种
	runSample(t, 2, 2, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 571)
}