package p3881

import "testing"

func runSample(t *testing.T, n int, pos int, k int, expect int) {
	res := countVisiblePeople(n, pos, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, pos, k := 3, 1, 0
	expect := 2
	runSample(t, n, pos, k, expect)
}

func TestSample2(t *testing.T) {
	n, pos, k := 3, 2, 1
	expect := 4
	runSample(t, n, pos, k, expect)
}

func TestSample3(t *testing.T) {
	n, pos, k := 1, 0, 0
	expect := 2
	runSample(t, n, pos, k, expect)
}

func TestSample4(t *testing.T) {
	n, pos, k := 3, 1, 1
	expect := 4
	runSample(t, n, pos, k, expect)
}
