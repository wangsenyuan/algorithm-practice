package p3871

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := countCommas(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := int64(1002)
	expect := int64(3)
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := int64(9980000000)
	expect := int64(28938999003)
	runSample(t, n, expect)
}