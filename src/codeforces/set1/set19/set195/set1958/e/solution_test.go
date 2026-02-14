package main

import "testing"

func runSample(t *testing.T, n, k int, expectNil bool) {
	res := solve(n, k)
	if expectNil {
		if res != nil {
			t.Errorf("Expected nil for n=%d k=%d, got %v", n, k, res)
		}
		return
	}
	if res == nil {
		t.Errorf("Expected solution for n=%d k=%d, got nil", n, k)
		return
	}
	if len(res) != n {
		t.Errorf("Expected length %d, got %d", n, len(res))
	}
	// Check it's a permutation
	seen := make(map[int]bool)
	for _, v := range res {
		if v < 1 || v > n {
			t.Errorf("Value %d out of range [1,%d]", v, n)
		}
		if seen[v] {
			t.Errorf("Duplicate value %d", v)
		}
		seen[v] = true
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 4, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 2, false)
}

func TestK1(t *testing.T) {
	runSample(t, 2, 1, false)
}

func TestK2(t *testing.T) {
	runSample(t, 3, 2, false)
}

func TestImpossible(t *testing.T) {
	runSample(t, 2, 2, true) // n=2, k=2: need f₂=3, impossible
}
