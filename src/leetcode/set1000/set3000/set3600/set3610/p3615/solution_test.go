package p3615

import "testing"

func runSample(t *testing.T, n int, edges [][]int, label string, expect int) {
	res := maxLen(n, edges, label)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}}
	label := "aabab"
	expect := 3
	runSample(t, n, edges, label, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1}, {0, 2}}
	label := "abc"
	expect := 1
	runSample(t, n, edges, label, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	edges := [][]int{{0, 2}, {0, 3}, {3, 1}}
	label := "bbac"
	expect := 3
	runSample(t, n, edges, label, expect)
}
