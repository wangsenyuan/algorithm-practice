package p3776

import "testing"

func runSample(t *testing.T, balance []int, expect int64) {
	res := minMoves(balance)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	balance := []int{5, 1, -4}
	expect := int64(4)
	runSample(t, balance, expect)
}

func TestSample2(t *testing.T) {
	balance := []int{1, 2, -5, 2}
	expect := int64(6)
	runSample(t, balance, expect)
}

func TestSample3(t *testing.T) {
	balance := []int{-3, 2}
	expect := int64(-1)
	runSample(t, balance, expect)
}

func TestSample4(t *testing.T) {
	balance := []int{14, 12, 7, -16}
	expect := int64(16)
	runSample(t, balance, expect)
}
