package p3805

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := countPairs(words)
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"fusion", "layout"}
	expect := 1
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"ab", "aa", "za", "aa"}
	expect := 2
	runSample(t, words, expect)
}
