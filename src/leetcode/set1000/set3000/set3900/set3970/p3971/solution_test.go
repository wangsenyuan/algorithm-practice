package p3971

import "testing"

func runSample(t *testing.T, value []int, decay []int, m int, expect int) {
	res := maxTotalValue(value, decay, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	value := []int{6, 5, 4}
	decay := []int{2, 1, 1}
	m := 4
	expect := 19
	runSample(t, value, decay, m, expect)
}

func TestSample2(t *testing.T) {
	value := []int{7, 2, 2}
	decay := []int{3, 2, 1}
	m := 2
	expect := 11
	runSample(t, value, decay, m, expect)
}

func TestSample3(t *testing.T) {
	value := []int{4, 3}
	decay := []int{5, 4}
	m := 10
	expect := 7
	runSample(t, value, decay, m, expect)
}
