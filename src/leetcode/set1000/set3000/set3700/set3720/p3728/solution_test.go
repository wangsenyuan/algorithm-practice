package p3728

import "testing"

func runSample(t *testing.T, capacity []int, expect int64) {
	res := countStableSubarrays(capacity)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	capacity := []int{9, 3, 3, 3, 9}
	expect := int64(2)
	runSample(t, capacity, expect)
}

func TestSample2(t *testing.T) {
	capacity := []int{-4, 4, 0, 0, -8, -4}
	expect := int64(1)
	runSample(t, capacity, expect)
}
