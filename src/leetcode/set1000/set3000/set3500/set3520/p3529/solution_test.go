package p3529

import "testing"

func runSample(t *testing.T, a [][]byte, pattern string, expect int) {
	res := countCells(a, pattern)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := [][]byte{[]byte("aacc"), []byte("bbbc"), []byte("aaba"), []byte("caac"), []byte("aacc")}
	pattern := "abaca"
	expect := 1
	runSample(t, grid, pattern, expect)
}
