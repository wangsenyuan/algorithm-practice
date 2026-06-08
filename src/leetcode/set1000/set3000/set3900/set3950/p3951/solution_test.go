package p3951

import "testing"

func runSample(t *testing.T, n int, brightness int, intervals [][]int, expect int) {
	res := minEnergy(n, brightness, intervals)

	if int64(expect) != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	brightness := 5
	intervals := [][]int{{6, 12}}
	expect := 14
	runSample(t, n, brightness, intervals, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	brightness := 1
	intervals := [][]int{{0, 0}, {2, 2}}
	expect := 2
	runSample(t, n, brightness, intervals, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	brightness := 2
	intervals := [][]int{{1, 3}, {2, 4}}
	expect := 4
	runSample(t, n, brightness, intervals, expect)
}
