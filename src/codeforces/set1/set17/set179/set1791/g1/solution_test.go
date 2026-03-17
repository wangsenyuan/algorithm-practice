package main

import "testing"

func runSample(t *testing.T, a []int, coins int, expect int) {
	res := solve(len(a), a, coins)

	if res != expect {
		t.Errorf("expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	coins := 6
	expect := 2
	runSample(t, a, coins, expect)
}

func TestSample2(t *testing.T) {
	a := []int{100, 52, 13, 6, 9, 4, 100, 35}
	coins := 32
	expect := 2
	runSample(t, a, coins, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5}
	coins := 1
	expect := 0
	runSample(t, a, coins, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 3, 2, 1}
	coins := 5
	expect := 1
	runSample(t, a, coins, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 3, 1, 4, 1}
	coins := 9
	expect := 2
	runSample(t, a, coins, expect)
}

func TestSample6(t *testing.T) {
	a := []int{500000000, 400000000, 300000000, 200000000, 100000000}
	coins := 600000000
	expect := 2
	runSample(t, a, coins, expect)
}
