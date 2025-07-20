package p3620

import "testing"

func runSample(t *testing.T, edges [][]int, online []bool, k int64, expect int) {
	res := findMaxPathScore(edges, online, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	edges := [][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}}
	online := []bool{true, true, true, true}
	k := int64(10)
	expect := 3
	runSample(t, edges, online, k, expect)
}

func TestSample2(t *testing.T) {
	edges := [][]int{{0, 1, 7}, {1, 4, 5}, {0, 2, 6}, {2, 3, 6}, {3, 4, 2}, {2, 4, 6}}
	online := []bool{true, true, true, false, true}
	k := int64(12)
	expect := 6
	runSample(t, edges, online, k, expect)
}
