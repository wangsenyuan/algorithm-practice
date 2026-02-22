package p3848

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := isDigitorialPermutation(n)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 40558, true)
}
