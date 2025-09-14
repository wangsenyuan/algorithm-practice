package p3679

import "testing"

func runSample(t *testing.T, arrive []int, w int, m int, expect int) {
	res := minArrivalsToDiscard(arrive, w, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arrivals := []int{1, 2, 1, 3, 1}
	w := 4
	m := 2
	expect := 0
	runSample(t, arrivals, w, m, expect)
}

func TestSample2(t *testing.T) {
	arrivals := []int{1, 2, 3, 3, 3, 4}
	w := 3
	m := 2
	expect := 1
	runSample(t, arrivals, w, m, expect)
}

func TestSample3(t *testing.T) {
	arrivals := []int{10, 4, 3, 6, 4, 5, 6, 1, 4}
	w := 7
	m := 1
	expect := 2
	runSample(t, arrivals, w, m, expect)
}
