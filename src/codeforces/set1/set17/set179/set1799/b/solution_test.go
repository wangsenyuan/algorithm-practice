package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, a []int, possible bool) {
	t.Helper()
	b := slices.Clone(a)
	ok, ops := solve(append([]int(nil), a...))
	if ok != possible {
		t.Fatalf("Sample a=%v possible=%v ok=%v ops=%v", a, possible, ok, ops)
	}

	if !possible {
		return
	}

	if len(ops) > 30*len(a) {
		t.Fatalf("len(ops) > 30*len(a) = %d > %d", len(ops), 30*len(a))
	}

	for _, op := range ops {
		i, j := op[0]-1, op[1]-1
		b[i] = (b[i] + b[j] - 1) / b[j]
	}

	mn := slices.Min(b)
	mx := slices.Max(b)
	if mn != mx {
		t.Fatalf("mn = %d != mx = %d", mn, mx)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{100}, true)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 1, 1}, true)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{2, 1}, false)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{5, 5}, true)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{4, 3, 2}, true)
}

func TestSample6(t *testing.T) {
	runSample(t, []int{3, 3, 4, 4}, true)
}

func TestSample7(t *testing.T) {
	runSample(t, []int{2, 100}, true)
}

func TestSample8(t *testing.T) {
	runSample(t, []int{5, 3, 6, 7, 8}, true)
}

func TestSample9(t *testing.T) {
	runSample(t, []int{3, 3, 80, 3, 8, 3}, true)
}

func TestSample10(t *testing.T) {
	runSample(t, []int{19, 40, 19, 55}, true)
}
