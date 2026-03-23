package main

import "testing"

func simulate(seq string) int {
	top, bottom := 0, 1
	last := 1
	for i := 0; i < len(seq); i++ {
		if seq[i] == 'T' {
			top += bottom
			last = top
		} else {
			bottom += top
			last = bottom
		}
	}
	return last
}

func countMistakes(seq string) int {
	var res int
	for i := 1; i < len(seq); i++ {
		if seq[i] == seq[i-1] {
			res++
		}
	}
	return res
}

func runSample(t *testing.T, n int, r int, expectMistakes int, expectPossible bool) {
	t.Helper()
	m, s := solve(n, r)
	if !expectPossible {
		if m >= 0 {
			t.Fatalf("expect impossible, but got mistakes=%d seq=%s", m, s)
		}
		return
	}
	if m != expectMistakes {
		t.Fatalf("expect mistakes %d, but got %d (seq=%s)", expectMistakes, m, s)
	}
	if len(s) != n {
		t.Fatalf("expect sequence length %d, but got %d", n, len(s))
	}
	if s[0] != 'T' {
		t.Fatalf("sequence must start with T, got %s", s)
	}
	if got := simulate(s); got != r {
		t.Fatalf("expect result %d, but simulation got %d from %s", r, got, s)
	}
	if got := countMistakes(s); got != m {
		t.Fatalf("reported mistakes %d, but counted %d in %s", m, got, s)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 10, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 5, 0, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, 0, false)
}

func TestSingleOperation(t *testing.T) {
	runSample(t, 1, 1, 0, true)
}
