package p3652

import "testing"

func runSample(t *testing.T, prices []int, strategy []int, k int, expect int64) {
	res := maxProfit(prices, strategy, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := []int{4, 2, 8}
	strategy := []int{-1, 0, 1}
	k := 2
	expect := int64(10)
	runSample(t, prices, strategy, k, expect)
}

func TestSample2(t *testing.T) {
	prices := []int{5, 4, 3}
	strategy := []int{1, 1, 0}
	k := 2
	expect := int64(9)
	runSample(t, prices, strategy, k, expect)
}

func TestSample3(t *testing.T) {
	prices := []int{5, 8}
	strategy := []int{-1, -1}
	k := 2
	expect := int64(8)
	runSample(t, prices, strategy, k, expect)
}
