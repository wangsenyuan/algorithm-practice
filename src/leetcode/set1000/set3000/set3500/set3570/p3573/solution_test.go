package p3573

import "testing"

func runSample(t *testing.T, prices []int, k int, expect int) {
	res := maximumProfit(prices, k)
	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	prices := []int{1, 7, 9, 8, 2}
	k := 2
	expect := 14
	runSample(t, prices, k, expect)
}


func TestSample2(t *testing.T) {
	prices := []int{12,16,19,19,8,1,19,13,9}
	k := 3
	expect := 36
	runSample(t, prices, k, expect)
}


func TestSample3(t *testing.T) {
	prices := []int{6,11,1,5,3,15,8}
	k := 3
	// 5 + 4 + 12 = 21
	// 10 + 12 = 22
	expect := 22
	runSample(t, prices, k, expect)
}